package viewmodels

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func ListOfSpeakers_CurrentSpeaker(ctx context.Context, los *dsmodels.ListOfSpeakers) (*dsmodels.Speaker, error) {
	var currentSpeaker *dsmodels.Speaker
	for _, speaker := range los.SpeakerList {
		if Speaker_IsCurrent(&speaker) {
			speechState := speaker.SpeechState

			if speechState == "interposed_question" {
				currentSpeaker = &speaker
				break
			} else {
				currentSpeaker = &speaker
			}
		}
	}

	return currentSpeaker, nil
}

type SpeakerListItem struct {
	Name                 string
	Weight               int
	IsSpeaking           bool
	IsContribution       bool
	IsPointOfOrder       bool
	IsIntervention       bool
	IsInterposedQuestion bool
	IsForspeach          bool
	IsCounterspeach      bool
}

type ListOfSpeakersLists struct {
	CurrentSpeaker             *SpeakerListItem
	WaitingSpeakers            []SpeakerListItem
	CurrentInterposedQuestion  *SpeakerListItem
	WaitingInterposedQuestions []SpeakerListItem
}

func ListOfSpeakers_CategorizedLists(ctx context.Context, fetch *dsmodels.Fetch, losID int) (ListOfSpeakersLists, error) {
	lQ := fetch.ListOfSpeakers(losID)
	los, err := lQ.
		Preload(lQ.SpeakerList().MeetingUser().StructureLevelList()).
		Preload(lQ.SpeakerList().MeetingUser().User()).First(ctx)
	if err != nil {
		return ListOfSpeakersLists{}, fmt.Errorf("could not load speakers: %w", err)
	}

	settingVal, err := fetch.Meeting_ListOfSpeakersDefaultStructureLevelTime(los.MeetingID).Value(ctx)
	if err != nil {
		return ListOfSpeakersLists{}, fmt.Errorf("could not fetch default_structure_level_time: %w", err)
	}

	waitingSpeakers := []SpeakerListItem{}
	interposedQuestions := []SpeakerListItem{}
	var currentSpeaker *SpeakerListItem
	var currentInterposedQuestion *SpeakerListItem
	for _, speaker := range los.SpeakerList {
		name := ""
		if meetingUser, isSet := speaker.MeetingUser.Value(); isSet {
			user := meetingUser.User
			name = User_ShortName(user)

			if settingVal == 0 {
				if len(meetingUser.StructureLevelList) != 0 {
					structureLevelNames := []string{}
					for _, sl := range meetingUser.StructureLevelList {
						structureLevelNames = append(structureLevelNames, sl.Name)
					}

					name = fmt.Sprintf("%s (%s)", name, strings.Join(structureLevelNames, ", "))
				}
			} else if settingVal > 0 {
				slID, ok := speaker.StructureLevelListOfSpeakersID.Value()
				if ok {
					slData, err := fetch.StructureLevel(slID).First(ctx)
					if err == nil {
						name = fmt.Sprintf("%s (%s)", name, slData.Name)
					}
				}
			}
		}

		item := SpeakerListItem{
			Name:                 name,
			Weight:               speaker.Weight,
			IsPointOfOrder:       speaker.PointOfOrder,
			IsContribution:       speaker.SpeechState == "contribution",
			IsIntervention:       speaker.SpeechState == "intervention",
			IsInterposedQuestion: speaker.SpeechState == "interposed_question",
			IsForspeach:          speaker.SpeechState == "pro",
			IsCounterspeach:      speaker.SpeechState == "contra",
			IsSpeaking:           false,
		}

		if speaker.BeginTime == 0 && speaker.EndTime == 0 {
			if speaker.SpeechState == "interposed_question" {
				interposedQuestions = append(interposedQuestions, item)
			} else {
				waitingSpeakers = append(waitingSpeakers, item)
			}
		} else if speaker.EndTime == 0 {
			if speaker.PauseTime == 0 {
				item.IsSpeaking = true
			}

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

	return ListOfSpeakersLists{
		CurrentSpeaker:             currentSpeaker,
		WaitingSpeakers:            waitingSpeakers,
		CurrentInterposedQuestion:  currentInterposedQuestion,
		WaitingInterposedQuestions: interposedQuestions,
	}, nil
}
