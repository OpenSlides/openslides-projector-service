package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	projectorHttp "github.com/OpenSlides/openslides-projector-service/pkg/http"
)

type config struct {
	Bind                 string `env:"BIND" envDefault:":9051"`
	Development          bool   `env:"OPENSLIDES_DEVELOPMENT" envDefault:"false"`
	PostgresHost         string `env:"DATABASE_HOST" envDefault:"localhost"`
	PostgresPort         string `env:"DATABASE_PORT" envDefault:"5432"`
	PostgresDatabase     string `env:"DATABASE_NAME" envDefault:"openslides"`
	PostgresUser         string `env:"DATABASE_USER" envDefault:"openslides"`
	PostgresPasswordFile string `env:"DATABASE_PASSWORD_FILE" envDefault:"/run/secrets/postgres_password"`
	MessageBusHost       string `env:"MESSAGE_BUS_HOST" envDetault:"localhost"`
	MessageBusPort       string `env:"MESSAGE_BUS_PORT" envDetault:"6379"`
	RestricterUrl        string `env:"RESTRICTER_URL" envDetault:"http://autoupdate:9012/internal/autoupdate"`
	PublicAccessOnly     bool   `env:"OPENSLIDES_PUBLIC_ACCESS_ONLY" envDefault:"false"`
}

func main() {
	var cfg config
	err := env.Parse(&cfg)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	if err != nil {
		log.Err(err).Msg("parsing config")
	}

	if cfg.Development {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err := run(cfg); err != nil {
		log.Fatal().Err(err).Msg("Error during startup")
	}

	log.Info().Msg("Stopped")
}

func run(cfg config) error {
	ctx := context.Background()

	env := &environment.ForProduction{}
	dsFlow, err := datastore.NewFlowPostgres(env)
	if err != nil {
		return fmt.Errorf("connecting to datastore: %w", err)
	}

	vote := datastore.NewFlowVoteCount(env)

	var dataFlow flow.Flow = dsFlow
	if !cfg.PublicAccessOnly {
		dataFlow = flow.Combine(
			dsFlow,
			map[string]flow.Flow{"poll/live_votes": vote},
		)

		eventer := func() (<-chan time.Time, func() bool) {
			timer := time.NewTimer(time.Second)
			return timer.C, timer.Stop
		}

		go vote.Connect(ctx, eventer, func(err error) {})
	}

	ds, err := getDatabase(cfg, dataFlow)
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}

	serverMux := http.NewServeMux()
	projectorHttp.New(ctx, projectorHttp.ProjectorConfig{
		RestricterUrl: cfg.RestricterUrl,
	}, serverMux, ds, dsFlow)
	fileHandler := http.StripPrefix("/system/projector/static/", http.FileServer(http.Dir("static")))
	serverMux.Handle("/system/projector/static/", fileHandler)

	log.Info().Msgf("Starting server on %s", cfg.Bind)
	srv := &http.Server{
		Addr:        cfg.Bind,
		Handler:     serverMux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Could not listen and serve")
	}

	return nil
}

func getDatabase(cfg config, dsFlow flow.Flow) (*database.Datastore, error) {
	password, err := parseSecretsFile(cfg.PostgresPasswordFile)
	if err != nil {
		if cfg.Development {
			password = "openslides"
		} else {
			return nil, fmt.Errorf("reading password from secrets: %w", err)
		}
	}

	pgAddr := fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(cfg.PostgresUser),
		encodePostgresConfig(password),
		encodePostgresConfig(cfg.PostgresHost),
		encodePostgresConfig(cfg.PostgresPort),
		encodePostgresConfig(cfg.PostgresDatabase),
	)
	redisAddr := cfg.MessageBusHost + ":" + cfg.MessageBusPort

	ds, err := database.New(pgAddr, redisAddr, dsFlow)
	if err != nil {
		return nil, fmt.Errorf("creating datastore: %w", err)
	}

	return ds, nil
}

// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

// parseSecretsFile takes a relative path as its argument
// and returns its contents
func parseSecretsFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
