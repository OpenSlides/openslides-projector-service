package viewmodels

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func User_Name(u *dsmodels.User) string {
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

func User_ShortName(u *dsmodels.User) string {
	nameParts := []string{}
	if title := u.Title; title != "" {
		nameParts = append(nameParts, title)
	}

	name := User_Name(u)
	nameParts = append(nameParts, name)
	return strings.Join(nameParts, " ")
}

func User_MeetingUserMap(ctx context.Context, fetch *dsmodels.Fetch, meetingID int) (map[int]int, error) {
	meetingUserIDs, err := fetch.Meeting_MeetingUserIDs(meetingID).Value(ctx)
	if err != nil {
		return nil, err
	}

	muUMap := map[int]*int{}
	for _, mu := range meetingUserIDs {
		var userID int
		fetch.MeetingUser_UserID(mu).Lazy(&userID)
		muUMap[mu] = &userID
	}

	if err := fetch.Execute(ctx); err != nil {
		return nil, err
	}

	uMuMap := map[int]int{}
	for mu, u := range muUMap {
		if muUMap[mu] != nil && *muUMap[mu] != 0 {
			uMuMap[*u] = mu
		}
	}

	return uMuMap, nil
}
