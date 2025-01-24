package models

import (
	"fmt"
	"strings"
)

func (m *MeetingUser) FullName() string {
	if m.user == nil {
		return ""
	}

	name := m.user.ShortName()
	additional := []string{}

	if m.user.Pronoun != nil && *m.user.Pronoun != "" {
		additional = append(additional, *m.user.Pronoun)
	}

	if m.structureLevels != nil {
		for _, sl := range m.StructureLevels() {
			additional = append(additional, sl.Name)
		}
	}

	if m.Number != nil && *m.Number != "" {
		additional = append(additional, fmt.Sprintf("No. %s", *m.Number))
	}

	if len(additional) == 0 {
		return name
	}

	return fmt.Sprintf("%s (%s)", name, strings.Join(additional, " · "))
}
