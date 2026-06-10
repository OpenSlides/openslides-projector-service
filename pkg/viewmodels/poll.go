package viewmodels

import (
	"encoding/json"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func Poll_ShouldShowChart(poll dsmodels.Poll) bool {
	switch config := poll.Config.(type) {
	case *dsmodels.PollConfigApproval:
		return true
	case *dsmodels.PollConfigSelection:
		return config.DisplayChart != ""
	}

	return false
}

func Poll_OneHundredPercentBase(poll dsmodels.Poll, option *dsmodels.PollOption) decimal.Decimal {
	switch config := poll.Config.(type) {
	case *dsmodels.PollConfigApproval:
		return Poll_OneHundredPercentBaseApproval(poll, config)
	case *dsmodels.PollConfigRatingApproval:
		return Poll_OneHundredPercentBaseRatingApproval(poll, config, option)
	case *dsmodels.PollConfigRatingScore:
		return Poll_OneHundredPercentBaseRatingScore(poll, config)
	case *dsmodels.PollConfigSelection:
		return Poll_OneHundredPercentBaseSelection(poll, config)
	}

	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseApproval(poll dsmodels.Poll, config *dsmodels.PollConfigApproval) decimal.Decimal {
	var result struct {
		Yes          string  `json:"yes"`
		No           string  `json:"no"`
		Abstain      *string `json:"abstain,omitempty"`
		TotalBallots int     `json:"total_ballots"`
	}

	err := json.Unmarshal([]byte(poll.Result), &result)
	if err != nil {
		log.Err(err).Msg("could not parse a poll result")
		return decimal.Decimal{}
	}

	switch config.OnehundredPercentBase {
	case "yes_no":
		yes, _ := decimal.NewFromString(result.Yes)
		no, _ := decimal.NewFromString(result.No)
		return yes.Add(no)
	case "valid":
		yes, _ := decimal.NewFromString(result.Yes)
		no, _ := decimal.NewFromString(result.No)
		abstain := decimal.Zero
		if result.Abstain != nil {
			abstain, _ = decimal.NewFromString(*result.Abstain)
		}
		return abstain.Add(yes).Add(no)
	case "cast":
		return decimal.NewFromInt(int64(result.TotalBallots))
	}

	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseSelection(poll dsmodels.Poll, config *dsmodels.PollConfigSelection) decimal.Decimal {
	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseRatingApproval(poll dsmodels.Poll, config *dsmodels.PollConfigRatingApproval, option *dsmodels.PollOption) decimal.Decimal {
	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseRatingScore(poll dsmodels.Poll, config *dsmodels.PollConfigRatingScore) decimal.Decimal {
	return decimal.Decimal{}
}

/*
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
	entitledUserIDsMap := map[int]struct{}{}
	meetingUserMap := make(map[int]dsmodels.MeetingUser)

	if poll.EntitledUsersAtStop != nil {
		var entitledUsersAtStop []struct {
			UserID int `json:"user_id"`
		}
		if err := json.Unmarshal(poll.EntitledUsersAtStop, &entitledUsersAtStop); err != nil {
			return []int{}
		}

		for _, entry := range entitledUsersAtStop {
			entitledUserIDsMap[entry.UserID] = struct{}{}
		}

		for _, group := range poll.EntitledGroupList {
			for _, mu := range group.MeetingUserList {
				meetingUserMap[mu.UserID] = mu
			}
		}
	} else {
		for _, group := range poll.EntitledGroupList {
			for _, mu := range group.MeetingUserList {
				entitledUserIDsMap[mu.UserID] = struct{}{}
				meetingUserMap[mu.UserID] = mu
			}
		}
	}

	if nameOrderSetting == "" {
		nameOrderSetting = "last_name"
	}

	entitledUserIDs := slices.Collect(maps.Keys(entitledUserIDsMap))
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
*/
