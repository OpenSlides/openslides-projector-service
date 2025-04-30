package viewmodels

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func MeetingUser_FullName(ctx context.Context, mu *dsmodels.MeetingUser) (string, error) {
	name := User_ShortName(mu.User)
	additional := []string{}
	if mu.User.Pronoun != "" {
		additional = append(additional, mu.User.Pronoun)
	}

	for _, sl := range mu.StructureLevelList {
		additional = append(additional, sl.Name)
	}

	if mu.Number != "" {
		// TODO: Translation
		additional = append(additional, fmt.Sprintf("No. %s", mu.Number))
	}

	if len(additional) == 0 {
		return name, nil
	}

	return fmt.Sprintf("%s (%s)", name, strings.Join(additional, " · ")), nil
}

func MeetingUser_StructureLevelNames(ctx context.Context, mu *dsmodels.MeetingUser) (string, error) {
	structureLevelNames := []string{}
	for _, sl := range mu.StructureLevelList {
		structureLevelNames = append(structureLevelNames, sl.Name)
	}

	return strings.Join(structureLevelNames, ", "), nil
}
