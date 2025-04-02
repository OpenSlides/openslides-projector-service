package slide

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (interface{}, error) {
	projection := req.Projection

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

	speakerIDs, err := req.Fetch.ListOfSpeakers_SpeakerIDs(*losID).Value(ctx)
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
