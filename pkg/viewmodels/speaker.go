package viewmodels

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

func Speaker_IsCurrent(s *dsfetch.Speaker) bool {
	return s.BeginTime != 0 && s.EndTime == 0
}

func Speaker_FullName(ctx context.Context, speaker *dsfetch.Speaker) (*string, error) {
	if meetingUserRef, isSet := speaker.MeetingUser().Value(); isSet {
		meetingUser, err := meetingUserRef.Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load meeting user: %w", err)
		}

		user, err := meetingUser.User().Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load user: %w", err)
		}

		name := User_ShortName(&user)
		return &name, nil
	}

	return nil, nil
}

func Speaker_StructureLevelName(ctx context.Context, speaker *dsfetch.Speaker) (*string, error) {
	if sllosRef, isSet := speaker.StructureLevelListOfSpeakers().Value(); isSet {
		sllos, err := sllosRef.Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load los structure level: %w", err)
		}

		sl, err := sllos.StructureLevel().Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load structure level: %w", err)
		}

		return &sl.Name, nil
	}

	return nil, nil
}
