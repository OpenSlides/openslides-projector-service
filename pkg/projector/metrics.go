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
	projectors := 0
	pool.projectors.Range(func(id string, val *projector) bool {
		renderedProjections += len(val.Projections)
		listeners += len(val.listeners)
		projectors++
		return true
	})

	metrics := map[string]int{
		"projectors":          projectors,
		"renderedProjections": renderedProjections,
		"subscribers":         listeners,
		"dbListeners":         pool.db.NumDsListeners(),
	}

	if data, err := json.Marshal(metrics); err == nil {
		log.Info().Str("metric", string(data)).Msg("")
	}
}
