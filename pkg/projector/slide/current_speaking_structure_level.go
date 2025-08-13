package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentSpeakingStructureLevelSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

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

	l := req.Fetch.ListOfSpeakers(*losID)
	los, err := l.Preload(l.StructureLevelListOfSpeakersList().StructureLevel()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	for _, sllos := range los.StructureLevelListOfSpeakersList {
		if sllos.CurrentStartTime == 0 {
			continue
		}

		countdownTime := sllos.RemainingTime + float64(sllos.CurrentStartTime)

		return map[string]any{
			"ID":            sllos.StructureLevelID,
			"Name":          sllos.StructureLevel.Name,
			"Color":         sllos.StructureLevel.Color,
			"CountdownTime": countdownTime,
		}, nil
	}

	return nil, nil
}
