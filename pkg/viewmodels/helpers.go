package viewmodels

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

type WeightedListEntry struct {
	Name        string
	MeetingUser *dsmodels.MeetingUser
	Weight      int
}

func CalcWeightedListNames(list []WeightedListEntry) {
	for i, entry := range list {
		if entry.MeetingUser != nil {
			meetingUser := entry.MeetingUser
			user := meetingUser.User
			list[i].Name = User_ShortName(user)
			if len(meetingUser.StructureLevelList) != 0 {
				structureLevelNames := []string{}
				for _, sl := range meetingUser.StructureLevelList {
					structureLevelNames = append(structureLevelNames, sl.Name)
				}

				list[i].Name = fmt.Sprintf("%s (%s)", list[i].Name, strings.Join(structureLevelNames, ", "))
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Weight < list[j].Weight
	})
}

func GetContentObjectField[V any](ctx context.Context, fetch *dsmodels.Fetch, field string, fqid string) (V, error) {
	var result V

	dsKey, err := dskey.FromStringf("%s/%s", fqid, field)
	if err != nil {
		return result, err
	}

	keys, err := fetch.Get(ctx, dsKey)
	if err != nil {
		return result, fmt.Errorf("load los id: %w", err)
	}

	if val, ok := keys[dsKey]; !ok || len(val) == 0 {
		return result, fmt.Errorf("not found id: %w", err)
	}

	if err := json.Unmarshal(keys[dsKey], &result); err != nil {
		return result, fmt.Errorf("parse los id: %w", err)
	}

	return result, nil
}
