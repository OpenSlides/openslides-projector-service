package viewmodels

import (
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/i18n"
)

func MeetingUser_FullName(locale *i18n.ProjectorLocale, mu *dsmodels.MeetingUser) string {
	name := User_ShortName(mu.User)
	additional := []string{}
	if mu.User.Pronoun != "" {
		additional = append(additional, mu.User.Pronoun)
	}

	for _, sl := range mu.StructureLevelList {
		additional = append(additional, sl.Name)
	}

	if mu.Number != "" {
		no := locale.Get("No.")
		additional = append(additional, fmt.Sprintf("%s %s", no, mu.Number))
	}

	if len(additional) == 0 {
		return name
	}

	return fmt.Sprintf("%s (%s)", name, strings.Join(additional, " · "))
}

func MeetingUser_StructureLevelNames(mu *dsmodels.MeetingUser) (string, error) {
	structureLevelNames := []string{}
	for _, sl := range mu.StructureLevelList {
		structureLevelNames = append(structureLevelNames, sl.Name)
	}

	return strings.Join(structureLevelNames, ", "), nil
}
