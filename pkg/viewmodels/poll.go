package viewmodels

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/shopspring/decimal"
)

func Poll_OneHundredPercentBase(poll dsmodels.Poll, option *dsmodels.Option) decimal.Decimal {
	if len(poll.OptionIDs) == 1 && option == nil {
		option = &poll.OptionList[0]
	}

	if len(poll.OptionIDs) == 0 {
		return decimal.Decimal{}
	}

	// YN and YNA need an selected option
	if option == nil && (poll.OnehundredPercentBase == "YN" || poll.OnehundredPercentBase == "YNA") {
		return decimal.Decimal{}
	}

	switch poll.OnehundredPercentBase {
	case "Y":
		total := poll.Votesvalid
		if poll.GlobalOption != nil {
			if globalOption, isSet := poll.GlobalOption.Value(); isSet {
				abstain := globalOption.Abstain
				no := globalOption.No
				total = total.Sub(no).Sub(abstain)
			}
		}
		return total
	case "YN":
		yes := option.Yes
		no := option.No
		return yes.Add(no)
	case "YNA":
		yes := option.Yes
		no := option.No
		abstain := option.Abstain

		return yes.Add(no).Add(abstain)
	case "valid":
		valid := poll.Votesvalid
		return valid
	case "entitled":
		entitled, err := Poll_EntitledUsers(poll)
		if err != nil {
			return decimal.Decimal{}
		}

		return decimal.NewFromInt(int64(len(entitled)))
	case "entitled_present":
		entitled, err := Poll_EntitledUsers(poll)
		if err != nil {
			return decimal.Decimal{}
		}

		present := int64(0)
		for _, u := range entitled {
			if u.Present {
				present++
			}
		}
		return decimal.NewFromInt(present)
	case "cast":
		cast := poll.Votescast
		return cast
	}

	return decimal.Decimal{}
}

type EntitledUsersAtStop []struct {
	UserID  int  `json:"user_id"`
	Present bool `json:"present"`
}

func Poll_EntitledUsers(poll dsmodels.Poll) (EntitledUsersAtStop, error) {
	var users EntitledUsersAtStop
	if err := json.Unmarshal(poll.EntitledUsersAtStop, &users); err != nil {
		return nil, fmt.Errorf("parse los id: %w", err)
	}

	return users, nil
}

func Poll_EntitledUserIDsSorted(poll dsmodels.Poll, nameOrderSetting string) []int {
	var entitledUserIDs []int
	meetingUserMap := make(map[int]dsmodels.MeetingUser)

	if poll.EntitledUsersAtStop != nil {
		var entitledUsersAtStop []struct {
			UserID int `json:"user_id"`
		}
		if err := json.Unmarshal(poll.EntitledUsersAtStop, &entitledUsersAtStop); err != nil {
			for _, entry := range entitledUsersAtStop {
				entitledUserIDs = append(entitledUserIDs, entry.UserID)
			}
			return entitledUserIDs
		}

		for _, entry := range entitledUsersAtStop {
			entitledUserIDs = append(entitledUserIDs, entry.UserID)
		}

		for _, group := range poll.EntitledGroupList {
			for _, mu := range group.MeetingUserList {
				meetingUserMap[mu.UserID] = mu
			}
		}
	} else {
		for _, group := range poll.EntitledGroupList {
			for _, mu := range group.MeetingUserList {
				entitledUserIDs = append(entitledUserIDs, mu.UserID)
				meetingUserMap[mu.UserID] = mu
			}
		}
	}

	if nameOrderSetting == "" {
		nameOrderSetting = "last_name"
	}

	slices.SortFunc(entitledUserIDs, func(aID, bID int) int {
		muA, aExists := meetingUserMap[aID]
		muB, bExists := meetingUserMap[bID]

		if !aExists || !bExists {
			if !aExists && !bExists {
				return 0
			}
			if !aExists {
				return 1
			}
			return -1
		}

		slAName := ""
		if len(muA.StructureLevelList) > 0 {
			slAName = muA.StructureLevelList[0].Name
		}
		slBName := ""
		if len(muB.StructureLevelList) > 0 {
			slBName = muB.StructureLevelList[0].Name
		}

		if slAName != slBName {
			return strings.Compare(slAName, slBName)
		}

		userA := muA.User
		userB := muB.User

		if nameOrderSetting == "first_name" {
			firstNameA := strings.Trim(userA.Title+" "+userA.FirstName, " ")
			firstNameB := strings.Trim(userB.Title+" "+userB.FirstName, " ")
			if firstNameA != firstNameB {
				return strings.Compare(firstNameA, firstNameB)
			}
			return strings.Compare(userA.LastName, userB.LastName)
		} else {
			if userA.LastName != userB.LastName {
				return strings.Compare(userA.LastName, userB.LastName)
			}
			firstNameA := strings.Trim(userA.Title+" "+userA.FirstName, " ")
			firstNameB := strings.Trim(userB.Title+" "+userB.FirstName, " ")
			return strings.Compare(firstNameA, firstNameB)
		}
	})

	return entitledUserIDs
}
