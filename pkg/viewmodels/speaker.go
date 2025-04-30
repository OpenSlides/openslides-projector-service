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
