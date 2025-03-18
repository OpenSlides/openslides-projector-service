package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (interface{}, error) {
	projection := req.Projection

	referenceProjectorId, err := req.Fetch.Meeting_ReferenceProjectorID(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return "", fmt.Errorf("could not load reference projector id %w", err)
	}

	refProjections, err := req.Fetch.Projector_CurrentProjectionIDs(referenceProjectorId).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load reference projector: %w", err)
	}

	losID := 0
	for _, pID := range refProjections {
		content, err := req.Fetch.Projection_ContentObjectID(pID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load projection: %w", err)
		}

		losDsKey, err := dskey.FromString("%s/list_of_speakers_id", content)
		if err != nil {
			continue
		}

		keys, err := req.Fetch.Get(ctx, losDsKey)
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

	speakerIDs, err := req.Fetch.ListOfSpeakers_SpeakerIDs(losID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load speakers: %w", err)
	}

	type speakerListItem struct {
		Name   string
		Weight int
	}
	waitingSpeakers := []speakerListItem{}
	interposedQuestions := []speakerListItem{}
	var currentSpeaker *speakerListItem
	var currentInterposedQuestion *speakerListItem
	for _, sID := range speakerIDs {
		speaker, err := req.Fetch.Speaker(sID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load speaker: %w", err)
		}

		name := ""
		if meetingUserRef, isSet := speaker.MeetingUser().Value(); isSet {
			meetingUser, err := meetingUserRef.Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("could not load meeting user: %w", err)
			}

			user, err := meetingUser.User().Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("could not load user: %w", err)
			}

			name = viewmodels.User_ShortName(&user)
			if len(meetingUser.StructureLevelList()) != 0 {
				structureLevelNames := []string{}
				for _, slRef := range meetingUser.StructureLevelList() {
					sl, err := slRef.Value(ctx)
					if err != nil {
						return "", err
					}

					structureLevelNames = append(structureLevelNames, sl.Name)
				}

				name = fmt.Sprintf("%s (%s)", name, strings.Join(structureLevelNames, ", "))
			}
		}

		item := speakerListItem{
			Name:   name,
			Weight: speaker.Weight,
		}

		if (speaker.BeginTime == 0) && speaker.EndTime == 0 {
			if speaker.SpeechState == "interposed_question" {
				interposedQuestions = append(interposedQuestions, item)
			} else {
				waitingSpeakers = append(waitingSpeakers, item)
			}
		} else if speaker.EndTime == 0 {
			if speaker.SpeechState == "interposed_question" {
				currentInterposedQuestion = &item
			} else {
				currentSpeaker = &item
			}
		}
	}

	sort.Slice(waitingSpeakers, func(i, j int) bool {
		return waitingSpeakers[i].Weight < waitingSpeakers[j].Weight
	})

	sort.Slice(interposedQuestions, func(i, j int) bool {
		return interposedQuestions[i].Weight < interposedQuestions[j].Weight
	})

	stable := false
	if projection.Stable != nil {
		stable = *projection.Stable
	}

	return map[string]interface{}{
		"CurrentSpeaker":            currentSpeaker,
		"Speakers":                  waitingSpeakers,
		"InterposedQuestions":       interposedQuestions,
		"CurrentInterposedQuestion": currentInterposedQuestion,
		"Overlay":                   stable,
	}, nil
}
