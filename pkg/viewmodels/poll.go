package viewmodels

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/shopspring/decimal"
)

func Poll_ShouldShowChart(poll dsmodels.Poll) bool {
	switch config := poll.Config.(type) {
	case *dsmodels.PollConfigApproval:
		return true
	case *dsmodels.PollConfigSelection:
		return config.DisplayChart == "pie"
	}

	return false
}

/*
func Poll_OneHundredPercentBase(poll dsmodels.Poll, option *dsmodels.PollOption) decimal.Decimal {
	switch config := poll.Config.(type) {
	case *dsmodels.PollConfigRatingApproval:
		return Poll_OneHundredPercentBaseRatingApproval(poll, config, option)
	case *dsmodels.PollConfigRatingScore:
		return Poll_OneHundredPercentBaseRatingScore(poll, config)
	case *dsmodels.PollConfigSelection:
		return Poll_OneHundredPercentBaseSelection(poll, config)
	}

	return decimal.Decimal{}
}
*/

func Poll_OneHundredPercentBaseSelection(poll dsmodels.Poll, config *dsmodels.PollConfigSelection) decimal.Decimal {
	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseRatingApproval(poll dsmodels.Poll, config *dsmodels.PollConfigRatingApproval, option *dsmodels.PollOption) decimal.Decimal {
	return decimal.Decimal{}
}

func Poll_OneHundredPercentBaseRatingScore(poll dsmodels.Poll, config *dsmodels.PollConfigRatingScore) decimal.Decimal {
	return decimal.Decimal{}
}

type PollResult interface {
	VotesInvalid() int64
	VotesValid() int64
	VotesCast() int64
}

type PollResultApproval struct {
	Yes          decimal.Decimal `json:"yes"`
	No           decimal.Decimal `json:"no"`
	Abstain      decimal.Decimal `json:"abstain"`
	Invalid      int             `json:"invalid"`
	TotalBallots int             `json:"total_ballots"`
}

func (r *PollResultApproval) VotesInvalid() int64 {
	return int64(r.Invalid)
}

func (r *PollResultApproval) VotesValid() int64 {
	return int64(r.TotalBallots - r.Invalid)
}

func (r *PollResultApproval) VotesCast() int64 {
	return int64(r.TotalBallots)
}

func (r *PollResultApproval) OneHundredPercentBase(config *dsmodels.PollConfigApproval) decimal.Decimal {
	switch config.OnehundredPercentBase {
	case "yes_no":
		return r.Yes.Add(r.No)
	case "valid":
		return r.Yes.Add(r.No).Add(r.Abstain)
	}

	return genericOnehundredPercentBase(r, config.OnehundredPercentBase)
}

type PollResultSelection struct {
	Options      map[string]decimal.Decimal `json:"-"`
	Nota         decimal.Decimal            `json:"nota"`
	Abstain      decimal.Decimal            `json:"abstain"`
	Invalid      int                        `json:"invalid"`
	TotalBallots int                        `json:"total_ballots"`
}

func (p *PollResultSelection) UnmarshalJSON(data []byte) error {
	type PollResultSelection_ PollResultSelection

	var aux PollResultSelection_
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("decode PollResultSelection: %w", err)
	}
	*p = PollResultSelection(aux)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("decode PollResultSelection raw map: %w", err)
	}

	delete(raw, "nota")
	delete(raw, "abstain")
	delete(raw, "invalid")
	delete(raw, "total_ballots")

	p.Options = make(map[string]decimal.Decimal, len(raw))
	for key, value := range raw {
		var d decimal.Decimal
		if err := json.Unmarshal(value, &d); err != nil {
			return fmt.Errorf("decode PollResultSelection option %q: %w", key, err)
		}
		p.Options[key] = d
	}

	return nil
}

func (r *PollResultSelection) VotesInvalid() int64 {
	return int64(r.Invalid)
}

func (r *PollResultSelection) VotesValid() int64 {
	return int64(r.TotalBallots - r.Invalid)
}

func (r *PollResultSelection) VotesCast() int64 {
	return int64(r.TotalBallots)
}

func (r *PollResultSelection) OneHundredPercentBase(config *dsmodels.PollConfigSelection) decimal.Decimal {
	// TODO: Add missing bases
	switch config.OnehundredPercentBase {
	case "no_general":
	case "valid":
	}

	return genericOnehundredPercentBase(r, config.OnehundredPercentBase)
}

type PollResultRatingScore struct {
	Options      map[string]decimal.Decimal `json:",inline"`
	Abstain      decimal.Decimal            `json:"abstain"`
	Invalid      int                        `json:"invalid"`
	TotalBallots int                        `json:"total_ballots"`
}

func (p *PollResultRatingScore) UnmarshalJSON(data []byte) error {
	type PollResultRatingScore_ PollResultRatingScore

	var aux PollResultRatingScore_
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("decode PollResultRatingScore: %w", err)
	}
	*p = PollResultRatingScore(aux)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("decode PollResultRatingScore raw map: %w", err)
	}

	delete(raw, "abstain")
	delete(raw, "invalid")
	delete(raw, "total_ballots")

	p.Options = make(map[string]decimal.Decimal, len(raw))
	for key, value := range raw {
		var d decimal.Decimal
		if err := json.Unmarshal(value, &d); err != nil {
			return fmt.Errorf("decode PollResultRatingScore option %q: %w", key, err)
		}
		p.Options[key] = d
	}

	return nil
}

func (r *PollResultRatingScore) VotesInvalid() int64 {
	return int64(r.Invalid)
}

func (r *PollResultRatingScore) VotesValid() int64 {
	return int64(r.TotalBallots - r.Invalid)
}

func (r *PollResultRatingScore) VotesCast() int64 {
	return int64(r.TotalBallots)
}

func (r *PollResultRatingScore) OneHundredPercentBase(config *dsmodels.PollConfigRatingScore) decimal.Decimal {
	// TODO: Add missing bases
	switch config.OnehundredPercentBase {
	}

	return genericOnehundredPercentBase(r, config.OnehundredPercentBase)
}

type PollResultRatingApprovalOption struct {
	Yes     decimal.Decimal `json:"yes"`
	No      decimal.Decimal `json:"no"`
	Abstain decimal.Decimal `json:"abstain"`
}

type PollResultRatingApproval struct {
	Options      map[string]PollResultRatingApprovalOption `json:",inline"`
	Abstain      decimal.Decimal                           `json:"abstain"`
	Invalid      int                                       `json:"invalid"`
	TotalBallots int                                       `json:"total_ballots"`
}

func (p *PollResultRatingApproval) UnmarshalJSON(data []byte) error {
	type PollResultRatingApproval_ PollResultRatingApproval

	var aux PollResultRatingApproval_
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("decode PollResultRatingApproval: %w", err)
	}
	*p = PollResultRatingApproval(aux)

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("decode PollResultRatingApproval raw map: %w", err)
	}

	delete(raw, "abstain")
	delete(raw, "invalid")
	delete(raw, "total_ballots")

	p.Options = make(map[string]PollResultRatingApprovalOption, len(raw))
	for key, value := range raw {
		var opt PollResultRatingApprovalOption
		if err := json.Unmarshal(value, &opt); err != nil {
			return fmt.Errorf("decode PollResultRatingApproval option %q: %w", key, err)
		}
		p.Options[key] = opt
	}

	return nil
}

func (r *PollResultRatingApproval) VotesInvalid() int64 {
	return int64(r.Invalid)
}

func (r *PollResultRatingApproval) VotesValid() int64 {
	return int64(r.TotalBallots - r.Invalid)
}

func (r *PollResultRatingApproval) VotesCast() int64 {
	return int64(r.TotalBallots)
}

func (r *PollResultRatingApproval) OneHundredPercentBase(config *dsmodels.PollConfigRatingApproval, option *dsmodels.PollOption) decimal.Decimal {
	if option == nil {
		return decimal.Decimal{}
	}

	// TODO: Add missing bases
	opt := r.Options[strconv.Itoa(option.ID)]
	switch config.OnehundredPercentBase {
	case "yes_no":
		return opt.Yes.Add(opt.No)
	case "valid":
		return opt.Yes.Add(opt.No).Add(opt.Abstain)
	}

	return genericOnehundredPercentBase(r, config.OnehundredPercentBase)
}

func genericOnehundredPercentBase(r PollResult, base string) decimal.Decimal {
	switch base {
	case "cast":
		return decimal.NewFromInt(r.VotesCast())
	case "valid":
		return decimal.NewFromInt(r.VotesValid())
	}

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
