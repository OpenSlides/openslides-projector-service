package slide

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentStructureLevelListSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
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
	los, err := l.Preload(l.StructureLevelListOfSpeakersList().SpeakerList()).
		Preload(l.StructureLevelListOfSpeakersList().StructureLevel()).
		Preload(l.SpeakerList().StructureLevelListOfSpeakers().StructureLevel()).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	type structureLevelEntry struct {
		ID            int     `json:"id"`
		Name          string  `json:"name"`
		Color         string  `json:"color"`
		SpeechState   string  `json:"speech_state"`
		CountdownTime float64 `json:"remaining_time"`
		Running       bool    `json:"current_start_time"`
	}
	structureLevels := []structureLevelEntry{}
	for _, sllos := range los.StructureLevelListOfSpeakersList {
		totalTime := float64(sllos.InitialTime) + sllos.AdditionalTime
		if totalTime == sllos.RemainingTime && sllos.CurrentStartTime == 0 {
			foundSpeaker := false
			for _, speaker := range sllos.SpeakerList {
				if speaker.SpeechState == "interposed_question" || speaker.SpeechState == "intervention" {
					continue
				}

				foundSpeaker = true
				break
			}

			if !foundSpeaker {
				continue
			}
		}

		countdownRunning := sllos.CurrentStartTime != 0
		countdownTime := sllos.RemainingTime
		if countdownRunning {
			countdownTime += float64(sllos.CurrentStartTime)
		}
		structureLevels = append(structureLevels, structureLevelEntry{
			ID:            sllos.StructureLevelID,
			Name:          sllos.StructureLevel.Name,
			Color:         sllos.StructureLevel.Color,
			CountdownTime: countdownTime,
			Running:       countdownRunning,
		})
	}

	var interventionSpeakers []dsmodels.Speaker
	for _, speaker := range los.SpeakerList {
		if speaker.SpeechState == "intervention" && speaker.EndTime == 0 {
			interventionSpeakers = append(interventionSpeakers, speaker)
		}
	}

	if len(interventionSpeakers) > 0 {
		defaultInterventionTime, err := req.Fetch.Meeting_ListOfSpeakersInterventionTime(los.MeetingID).Value(ctx)
		name := "\nIntervention"
		if err != nil {
			return nil, fmt.Errorf("couldn not load intervention time %w", err)
		}
		interventionEntry := structureLevelEntry{
			Name:          name,
			CountdownTime: float64(defaultInterventionTime),
		}

		var currentInterventionSpeaker *dsmodels.Speaker
		for _, s := range interventionSpeakers {
			if s.BeginTime != 0 && s.EndTime == 0 {
				currentInterventionSpeaker = &s
			}
		}

		if currentInterventionSpeaker != nil {
			running := currentInterventionSpeaker.PauseTime == 0
			interventionEntry.Running = running

			if currentInterventionSpeaker.PauseTime == 0 {
				interventionEntry.CountdownTime = float64(currentInterventionSpeaker.BeginTime) + float64(defaultInterventionTime) + float64(currentInterventionSpeaker.TotalPause)
			} else {
				now := int(time.Now().Unix())
				elapsed := now - currentInterventionSpeaker.BeginTime - currentInterventionSpeaker.TotalPause
				remaining := float64(defaultInterventionTime) - float64(elapsed)
				interventionEntry.CountdownTime = remaining
			}
			if currentInterventionSpeaker.StructureLevelListOfSpeakers != nil {
				sllos, ok := currentInterventionSpeaker.StructureLevelListOfSpeakers.Value()
				if ok {
					interventionEntry.ID = sllos.StructureLevelID
					interventionEntry.Color = sllos.StructureLevel.Color
				}
			}
		}

		structureLevels = append(structureLevels, interventionEntry)
	}

	titleInfo, err := viewmodels.GetTitleInformationByContentObject(ctx, req.Fetch, los.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load los title info %w", err)
	}

	return map[string]any{
		"ContentTitle":    titleInfo,
		"StructureLevels": structureLevels,
	}, nil
}
