package viewmodels

import (
	"fmt"
	"sort"
	"strings"

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
