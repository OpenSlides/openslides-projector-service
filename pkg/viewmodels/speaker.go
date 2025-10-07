package viewmodels

import (
	"context"
	"time"

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
		return float64(speaker.BeginTime) + float64(interventionTime) + float64(speaker.TotalPause)
	} else {
		now := int(time.Now().Unix())
		elapsed := now - speaker.BeginTime - speaker.TotalPause
		return float64(interventionTime) - float64(elapsed)
	}
}
