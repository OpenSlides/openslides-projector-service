package viewmodels

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

func MeetingUser_FullName(ctx context.Context, mu *dsfetch.MeetingUser) (string, error) {
	user, err := mu.User().Value(ctx)
	if err != nil {
		return "", err
	}

	name := User_ShortName(&user)
	additional := []string{}
	if user.Pronoun != "" {
		additional = append(additional, user.Pronoun)
	}

	for _, slRef := range mu.StructureLevelList() {
		sl, err := slRef.Value(ctx)
		if err != nil {
			return "", err
		}

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

func MeetingUser_StructureLevelNames(ctx context.Context, mu *dsfetch.MeetingUser) (string, error) {
	structureLevelNames := []string{}
	for _, slRef := range mu.StructureLevelList() {
		sl, err := slRef.Value(ctx)
		if err != nil {
			return "", fmt.Errorf("could not load structure level: %w", err)
		}

		structureLevelNames = append(structureLevelNames, sl.Name)
	}

	return strings.Join(structureLevelNames, ", "), nil
}
