package viewmodels

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func Speaker_IsCurrent(s *dsmodels.Speaker) bool {
	return s.BeginTime != 0 && s.EndTime == 0
}

func Speaker_FullName(ctx context.Context, speaker *dsmodels.Speaker) (*string, error) {
	if meetingUser, isSet := speaker.MeetingUser.Value(); isSet {
		user := meetingUser.User
		name := User_ShortName(user)
		return &name, nil
	}

	return nil, nil
}

func Speaker_StructureLevelName(ctx context.Context, speaker *dsmodels.Speaker) (*string, error) {
	if sllos, isSet := speaker.StructureLevelListOfSpeakers.Value(); isSet {
		return &sllos.StructureLevel.Name, nil
	}

	return nil, nil
}

func Speaker_CalculateInterventionCountdownTime(speaker *dsmodels.Speaker, interventionTime int) float64 {
	if speaker == nil {
		return 0
	}

	if speaker.PauseTime == 0 {
		return float64(interventionTime) + Speaker_CalculateElapsedTime(speaker)
	} else {
		return float64(interventionTime) - Speaker_CalculateElapsedTime(speaker)
	}
}

func Speaker_CalculateElapsedTime(speaker *dsmodels.Speaker) float64 {
	if speaker == nil || speaker.BeginTime == 0 {
		return 0
	}

	if speaker.PauseTime == 0 {
		return float64(speaker.BeginTime + speaker.TotalPause)
	} else {
		elapsed := speaker.PauseTime - speaker.BeginTime - speaker.TotalPause
		return float64(elapsed)
	}
}
