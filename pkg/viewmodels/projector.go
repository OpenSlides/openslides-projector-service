package viewmodels

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func Projector_ListOfSpeakersID(ctx context.Context, fetch *dsmodels.Fetch, projectorID int) (*int, error) {
	projections, err := fetch.Projector_CurrentProjectionIDs(projectorID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load reference projector: %w", err)
	}

	for _, pID := range projections {
		content, err := fetch.Projection_ContentObjectID(pID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load projection: %w", err)
		}

		losID, err := GetContentObjectField[int](ctx, fetch, "list_of_speakers_id", content)
		if err != nil {
			return nil, err
		}

		if losID != nil {
			return losID, nil
		}
	}

	return nil, nil
}
