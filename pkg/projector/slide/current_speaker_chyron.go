package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

type currentSpeakerChyronSlideOptions struct {
	ChyronType string `json:"chyron_type"`
	AgendaItem bool   `json:"agenda_item"`
}

func CurrentSpeakerChyronSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
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

	losQ := req.Fetch.ListOfSpeakers(*losID)
	los, err := losQ.
		Preload(
			losQ.SpeakerList().
				MeetingUser().
				StructureLevelList(),
		).
		Preload(
			losQ.SpeakerList().
				StructureLevelListOfSpeakers().
				StructureLevel(),
		).
		Preload(
			losQ.SpeakerList().
				MeetingUser().
				User(),
		).
		First(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers: %w", err)
	}

	var options currentSpeakerChyronSlideOptions
	if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
		return nil, fmt.Errorf("could not parse slide options: %w", err)
	}

	currentSpeaker, err := viewmodels.ListOfSpeakers_CurrentSpeaker(ctx, &los)
	if err != nil {
		return nil, fmt.Errorf("could not get current speaker: %w", err)
	}

	slideSpeakerName := ""
	slideStructureLevel := ""
	slideAgendaItem := ""
	if currentSpeaker != nil {
		speakerName, err := viewmodels.Speaker_FullName(ctx, currentSpeaker)
		if err != nil {
			return nil, fmt.Errorf("could not get speaker name: %w", err)
		}

		if speakerName != nil {
			slideSpeakerName = *speakerName

			structureLevelDefaultTime, err := req.Fetch.Meeting_ListOfSpeakersDefaultStructureLevelTime(los.MeetingID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("could not load ListOfSpeakersDefaultStructureLevelTime: %w", err)
			}

			if structureLevelDefaultTime > 0 {
				structureLevel, err := viewmodels.Speaker_StructureLevelName(ctx, currentSpeaker)
				if err != nil {
					return nil, fmt.Errorf("could not get speaker structurelevels: %w", err)
				}

				if structureLevel != nil {
					slideStructureLevel = *structureLevel
				}
			} else {
				if meetingUser, isSet := currentSpeaker.MeetingUser.Value(); isSet {
					structureLevels, err := viewmodels.MeetingUser_StructureLevelNames(ctx, &meetingUser)
					if err != nil {
						return nil, fmt.Errorf("could not load structure levels: %w", err)
					}

					slideStructureLevel = structureLevels
				}
			}

			if options.ChyronType == "new" && slideStructureLevel != "" {
				slideSpeakerName = fmt.Sprintf("%s, %s", slideSpeakerName, slideStructureLevel)
			}
		}
	}

	co, err := viewmodels.GetTitleInformationByContentObject(ctx, req.Fetch, los.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch coTitle information")
	}
	switch co.Collection {
	case "motion", "topic":
		slideAgendaItem = strings.Join(func() []string {
			parts := []string{}
			for _, s := range []string{co.AgendaItemNumber, co.Number, co.Title} {
				if strings.TrimSpace(s) != "" {
					parts = append(parts, s)
				}
			}
			return parts
		}(), " · ")
	default:
		slideAgendaItem = co.Title
	}

	return map[string]any{
		"Options":        options,
		"SpeakerName":    slideSpeakerName,
		"StructureLevel": slideStructureLevel,
		"AgendaItem":     slideAgendaItem,
	}, nil
}
