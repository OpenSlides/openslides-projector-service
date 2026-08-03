package viewmodels

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/i18n"
)

func Option_OptionLabel(ctx context.Context, fetch *dsmodels.Fetch, locale *i18n.ProjectorLocale, option *dsmodels.PollOption) (string, error) {
	if option.Text != "" {
		return option.Text, nil
	} else if muID, isSet := option.MeetingUserID.Value(); isSet {
		muQ := fetch.MeetingUser(muID)
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
	}

	return "", nil
}
