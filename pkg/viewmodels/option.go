package viewmodels

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/i18n"
)

func Option_OptionLabel(ctx context.Context, fetch *dsmodels.Fetch, locale *i18n.ProjectorLocale, option *dsmodels.Option, userMap map[int]int) (string, error) {
	if option.Text != "" {
		return option.Text, nil
	} else if !option.ContentObjectID.Null() {
		contentObjectID, _ := option.ContentObjectID.Value()
		if strings.HasPrefix(contentObjectID, "user/") {
			if userMap == nil {
				var err error
				userMap, err = User_MeetingUserMap(ctx, fetch, option.MeetingID)
				if err != nil {
					return "", err
				}
			}

			idStr := contentObjectID[strings.Index(contentObjectID, "/")+1:]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return "", fmt.Errorf("could not parse poll option fqid: %w", err)
			}

			muQ := fetch.MeetingUser(userMap[id])
			mu, err := muQ.Preload(muQ.User()).Preload(muQ.StructureLevelList()).First(ctx)
			if err != nil {
				return "", fmt.Errorf("could not fetch poll option meeting user: %w", err)
			}

			slName := ""
			if len(mu.StructureLevelIDs) > 0 {
				slNames := make([]string, 0, len(mu.StructureLevelList))
				for _, sl := range mu.StructureLevelList {
					slNames = append(slNames, sl.Name)
				}
				slName = fmt.Sprintf(" (%s)", strings.Join(slNames, ", "))
			}
			return fmt.Sprintf("%s %s%s", mu.User.FirstName, mu.User.LastName, slName), nil
		} else if strings.HasPrefix(contentObjectID, "poll_candidate_list/") {
			return locale.Get("Confirmation of the nomination list"), nil
		}
	}

	return "", nil
}
