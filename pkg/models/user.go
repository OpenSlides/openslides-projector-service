package models

import (
	"fmt"
	"strings"
)

func (m *User) Name() string {
  nameParts := []string{}
  if firstName := m.FirstName; firstName != nil {
		nameParts = append(nameParts, *firstName)
	}

	if lastName := m.LastName; lastName != nil {
		nameParts = append(nameParts, *lastName)
	}

  if len(nameParts) == 0 {
    return fmt.Sprintf("User %d", m.ID)
  }

  return strings.Join(nameParts, " ")
}

func (m *User) ShortName() string {
  nameParts := []string{}
  if title := m.Title; title != nil {
		nameParts = append(nameParts, *title)
	}

	nameParts = append(nameParts, m.Name())
  return strings.Join(nameParts, " ")
}
