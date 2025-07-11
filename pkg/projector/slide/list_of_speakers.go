package slide

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func ListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no list of speakers id provided for slide")
	}

	losID := *req.ContentObjectID
	if strings.HasPrefix(req.Projection.ContentObjectID, "meeting") {
		referenceProjectorId, err := req.Fetch.Meeting_ReferenceProjectorID(*req.ContentObjectID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load reference projector id %w", err)
		}

		currentLosID, err := viewmodels.Projector_ListOfSpeakersID(ctx, req.Fetch, referenceProjectorId)
		if err != nil {
			return nil, fmt.Errorf("could not load list of speakers id %w", err)
		}

		if currentLosID == nil {
			return nil, nil
		} else {
			losID = *currentLosID
		}
	}

	los, err := req.Fetch.ListOfSpeakers(losID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	titleInfo, err := viewmodels.GetTitleInformationByContentObject(ctx, req.Fetch, los.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load los title info %w", err)
	}

	speakers, err := viewmodels.ListOfSpeakers_CategorizedLists(ctx, req.Fetch, los.ID)
	if err != nil {
		return nil, fmt.Errorf("could not categorize speakers %w", err)
	}

	if req.Projection.Stable &&
		len(speakers.WaitingInterposedQuestions) == 0 &&
		len(speakers.WaitingSpeakers) == 0 &&
		speakers.CurrentSpeaker == nil &&
		speakers.CurrentInterposedQuestion == nil {
		return nil, nil
	}

	return map[string]any{
		"_template":    "list_of_speakers",
		"LoS":          los,
		"Speakers":     speakers,
		"ContentTitle": titleInfo,
		"Overlay":      req.Projection.Stable,
	}, nil
}
