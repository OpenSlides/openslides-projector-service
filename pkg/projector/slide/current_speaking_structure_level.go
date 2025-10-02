package slide

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
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

	var currentSpeaker *dsmodels.Speaker
	for _, speaker := range los.SpeakerList {
		if speaker.BeginTime != 0 && speaker.EndTime == 0 {
			currentSpeaker = &speaker
			break
		}
	}

	if currentSpeaker == nil {
		return nil, nil
	}

	var currentSpeakerInfo speakerInfo
	currentSpeakerInfo.Running = currentSpeaker.PauseTime == 0

	if currentSpeaker.SpeechState == "intervention" {
		defaultInterventionTime, err := req.Fetch.Meeting_ListOfSpeakersInterventionTime(los.MeetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load intervention time: %w", err)
		}

		if currentSpeaker.PauseTime == 0 {
			currentSpeakerInfo.CountdownTime = float64(currentSpeaker.BeginTime) + float64(defaultInterventionTime) + float64(currentSpeaker.TotalPause)
		} else {
			now := int(time.Now().Unix())
			elapsed := now - currentSpeaker.BeginTime - currentSpeaker.TotalPause
			remaining := float64(defaultInterventionTime) - float64(elapsed)
			currentSpeakerInfo.CountdownTime = remaining
		}
		if currentSpeaker.StructureLevelListOfSpeakers != nil {
			sllos, ok := currentSpeaker.StructureLevelListOfSpeakers.Value()
			if ok {
				currentSpeakerInfo.ID = sllos.StructureLevelID
				currentSpeakerInfo.Name = sllos.StructureLevel.Name
				currentSpeakerInfo.Color = sllos.StructureLevel.Color
			}
		}

		currentSpeakerInfo.Name += "\nIntervention"
	} else {
		sllos, ok := currentSpeaker.StructureLevelListOfSpeakers.Value()
		if ok {
			currentSpeakerInfo.ID = sllos.StructureLevelID
			currentSpeakerInfo.Name = sllos.StructureLevel.Name
			currentSpeakerInfo.Color = sllos.StructureLevel.Color
			currentSpeakerInfo.CountdownTime = sllos.RemainingTime + float64(sllos.CurrentStartTime)
		}
	}
	return map[string]any{
		"SpeakerInfo": currentSpeakerInfo,
	}, nil
}

type speakerInfo struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Color         string  `json:"color"`
	CountdownTime float64 `json:"remaining_time"`
	Running       bool    `json:"running"`
}
