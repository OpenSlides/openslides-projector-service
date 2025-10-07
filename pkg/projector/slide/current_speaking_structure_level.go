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
	los, err := l.Preload(l.SpeakerList().StructureLevelListOfSpeakers().StructureLevel()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	currentSpeaker, err := viewmodels.ListOfSpeakers_CurrentSpeaker(ctx, &los)
	if err != nil {
		return nil, fmt.Errorf("could not fetch current speaker %w", err)
	}

	if currentSpeaker == nil {
		return nil, nil
	}

	type speakerInfo struct {
		ID            int
		Name          string
		Color         string
		CountdownTime float64
		Running       bool
		Intervention  bool
	}

	var currentSpeakerInfo speakerInfo
	currentSpeakerInfo.Running = currentSpeaker.PauseTime == 0
	sllos, ok := currentSpeaker.StructureLevelListOfSpeakers.Value()
	if ok {
		currentSpeakerInfo.ID = sllos.StructureLevelID
		currentSpeakerInfo.Name = sllos.StructureLevel.Name
		currentSpeakerInfo.Color = sllos.StructureLevel.Color
		currentSpeakerInfo.CountdownTime = sllos.RemainingTime + float64(sllos.CurrentStartTime)
	}

	if currentSpeaker.SpeechState != "intervention" && sllos.StructureLevelID == 0 {
		return nil, nil
	}

	if currentSpeaker.SpeechState == "intervention" {
		currentSpeakerInfo.Intervention = true
		defaultInterventionTime, err := req.Fetch.Meeting_ListOfSpeakersInterventionTime(los.MeetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load intervention time: %w", err)
		}

		currentSpeakerInfo.CountdownTime = viewmodels.Speaker_CalculateInterventionCountdownTime(currentSpeaker, defaultInterventionTime)
	}

	return map[string]any{
		"SpeakerInfo": currentSpeakerInfo,
	}, nil
}
