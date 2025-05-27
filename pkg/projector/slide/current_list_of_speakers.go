package slide

import (
	"context"
	"fmt"
	"sort"
	"strings"

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

	lQ := req.Fetch.ListOfSpeakers(*losID)
	los, err := lQ.
		Preload(lQ.SpeakerList().MeetingUser().StructureLevelList()).
		Preload(lQ.SpeakerList().MeetingUser().User()).First(ctx)
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
	for _, speaker := range los.SpeakerList {
		name := ""
		if meetingUser, isSet := speaker.MeetingUser.Value(); isSet {
			user := meetingUser.User
			name = viewmodels.User_ShortName(user)
			if len(meetingUser.StructureLevelList) != 0 {
				structureLevelNames := []string{}
				for _, sl := range meetingUser.StructureLevelList {
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

	return map[string]any{
		"CurrentSpeaker":            currentSpeaker,
		"Speakers":                  waitingSpeakers,
		"InterposedQuestions":       interposedQuestions,
		"CurrentInterposedQuestion": currentInterposedQuestion,
		"Overlay":                   req.Projection.Stable,
	}, nil
}
