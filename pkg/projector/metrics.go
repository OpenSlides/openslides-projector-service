package projector

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

// Logs metrics in a given duration
//
// Blocks until the context is done.
func MetricLoop(ctx context.Context, d time.Duration, pool *ProjectorPool) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			logMetricMessage(pool)
		}
	}
}

func logMetricMessage(pool *ProjectorPool) {
	renderedProjections := 0
	listeners := 0
	for _, projector := range pool.projectors {
		renderedProjections += len(projector.Projections)
		listeners += len(projector.listeners)
	}

	metrics := map[string]int{
		"projectors":          len(pool.projectors),
		"renderedProjections": renderedProjections,
		"subscribers":         listeners,
		"dbListeners":         pool.db.NumDsListeners(),
	}

	if data, err := json.Marshal(metrics); err != nil {
		log.Info().Msg(string(data))
	}
}
