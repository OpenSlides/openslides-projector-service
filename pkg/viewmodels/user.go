package viewmodels

import (
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

func User_Name(u *dsfetch.User) string {
	nameParts := []string{}
	if firstName := u.FirstName; firstName != "" {
		nameParts = append(nameParts, firstName)
	}

	if lastName := u.LastName; lastName != "" {
		nameParts = append(nameParts, lastName)
	}

	if len(nameParts) == 0 {
		return fmt.Sprintf("User %d", u.ID)
	}

	return strings.Join(nameParts, " ")
}

func User_ShortName(u *dsfetch.User) string {
	nameParts := []string{}
	if title := u.Title; title != "" {
		nameParts = append(nameParts, title)
	}

	name := User_Name(u)
	nameParts = append(nameParts, name)
	return strings.Join(nameParts, " ")
}
