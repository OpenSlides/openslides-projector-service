package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	referenceProjectorId, err := req.Fetch.Meeting_ReferenceProjectorID(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load reference projector id %w", err)
	}

	losID, err := viewmodels.Projector_ListOfSpeakersID(ctx, req.Fetch, referenceProjectorId)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers id %w", err)
	}

	if losID == nil {
		return nil, nil
	}

	speakers, err := viewmodels.ListOfSpeakers_CategorizedLists(ctx, req.Fetch, *losID)
	if err != nil {
		return nil, fmt.Errorf("could not categorize speakers %w", err)
	}

	if len(speakers.WaitingInterposedQuestions) == 0 &&
		len(speakers.WaitingSpeakers) == 0 &&
		speakers.CurrentSpeaker == nil &&
		speakers.CurrentInterposedQuestion == nil {
		return nil, nil
	}

	return map[string]any{
		"CurrentSpeaker":            speakers.CurrentSpeaker,
		"Speakers":                  speakers.WaitingSpeakers,
		"InterposedQuestions":       speakers.WaitingInterposedQuestions,
		"CurrentInterposedQuestion": speakers.CurrentInterposedQuestion,
		"Overlay":                   req.Projection.Stable,
	}, nil
}
