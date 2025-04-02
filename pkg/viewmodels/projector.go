package viewmodels

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
)

func Projector_ListOfSpeakersID(ctx context.Context, fetch *dsfetch.Fetch, projectorID int) (*int, error) {
	projections, err := fetch.Projector_CurrentProjectionIDs(projectorID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load reference projector: %w", err)
	}

	losID := 0
	for _, pID := range projections {
		content, err := fetch.Projection_ContentObjectID(pID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load projection: %w", err)
		}

		losDsKey, err := dskey.FromString("%s/list_of_speakers_id", content)
		if err != nil {
			continue
		}

		keys, err := fetch.Get(ctx, losDsKey)
		if err != nil {
			return nil, fmt.Errorf("load los id: %w", err)
		}

		if val, ok := keys[losDsKey]; !ok || len(val) == 0 {
			continue
		}

		if err := json.Unmarshal(keys[losDsKey], &losID); err != nil {
			return nil, fmt.Errorf("parse los id: %w", err)
		}
	}

	if losID == 0 {
		return nil, nil
	}

	return &losID, nil
}
