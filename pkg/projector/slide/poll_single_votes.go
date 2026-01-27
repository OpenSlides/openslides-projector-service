package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSingleVotesSlideVoteEntry struct {
	Value     string
	Present   bool
	FirstName string
	LastName  string
}

type pollSingleVotesSlideVoteEntryGroup struct {
	Title string
	Votes []*pollSingleVotesSlideVoteEntry
}

func (e *pollSingleVotesSlideVoteEntryGroup) TotalYes() int {
	sum := 0
	for _, v := range e.Votes {
		if v.Value == "Y" {
			sum += 1
		}
	}

	return sum
}

func (e *pollSingleVotesSlideVoteEntryGroup) TotalNo() int {
	sum := 0
	for _, v := range e.Votes {
		if v.Value == "N" {
			sum += 1
		}
	}

	return sum
}

func (e *pollSingleVotesSlideVoteEntryGroup) TotalAbstain() int {
	sum := 0
	for _, v := range e.Votes {
		if v.Value == "A" {
			sum += 1
		}
	}

	return sum
}

type pollSingleVotesSlideData struct {
	TotalVotesvalid decimal.Decimal
	PercVotesvalid  decimal.Decimal
	Options         []*pollSingleVotesSlideOption
	GroupedVotes    []*pollSingleVotesSlideVoteEntryGroup
}

type pollSingleVotesSlideOption struct {
	TotalYes     decimal.Decimal
	TotalNo      decimal.Decimal
	TotalAbstain decimal.Decimal
	PercYes      decimal.Decimal
	PercNo       decimal.Decimal
	PercAbstain  decimal.Decimal
}

func pollSingleVotesSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pQ := req.Fetch.Poll()
	poll, err := req.Fetch.Poll(*req.ContentObjectID).
		Preload(pQ.OptionList().VoteList()).
		Preload(pQ.GlobalOption().VoteList()).
		Preload(pQ.EntitledGroupList().MeetingUserList().User()).
		Preload(pQ.EntitledGroupList().MeetingUserList().VoteDelegatedTo().User()).
		Preload(pQ.EntitledGroupList().MeetingUserList().VoteDelegatedTo().User().IsPresentInMeetingList()).
		Preload(pQ.EntitledGroupList().MeetingUserList().User().IsPresentInMeetingList()).
		Preload(pQ.EntitledGroupList().MeetingUserList().StructureLevelList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	var maxColumns int
	var nameOrderString string
	req.Fetch.Meeting_MotionPollProjectionMaxColumns(poll.MeetingID).Lazy(&maxColumns)
	req.Fetch.Meeting_MotionPollProjectionNameOrderFirst(poll.MeetingID).Lazy(&nameOrderString)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not load meeting settings: %w", err)
	}

	if nameOrderString == "" {
		nameOrderString = "last_name"
	}

	voteMap, err := mapUsersToVote(&poll)
	if err != nil {
		return nil, fmt.Errorf("mapping users to vote: %w", err)
	}

	meetingUserMap := map[int]dsmodels.MeetingUser{}
	for _, group := range poll.EntitledGroupList {
		for _, mu := range group.MeetingUserList {
			meetingUserMap[mu.UserID] = mu
		}
	}

	slices.SortFunc(poll.OptionList, func(a dsmodels.Option, b dsmodels.Option) int {
		return a.Weight - b.Weight
	})

	optionIndexMap := map[string]int{}
	for idx, option := range poll.OptionList {
		optionIndexMap[strconv.Itoa(option.ID)] = idx
	}

	slideData := pollSingleVotesSlideData{}
	voteEntryGroupsMap := map[int]*pollSingleVotesSlideVoteEntryGroup{}
	entitledUsers := viewmodels.Poll_EntitledUserIDsSorted(poll, nameOrderString)
	for _, userID := range entitledUsers {
		mu, exists := meetingUserMap[userID]
		if !exists {
			continue
		}

		structureLevel := &dsmodels.StructureLevel{
			ID:   0,
			Name: "",
		}
		if len(mu.StructureLevelList) > 0 {
			structureLevel = &mu.StructureLevelList[0]
		}

		if _, ok := voteEntryGroupsMap[structureLevel.ID]; !ok {
			voteEntryGroupsMap[structureLevel.ID] = &pollSingleVotesSlideVoteEntryGroup{
				Title: structureLevel.Name,
				Votes: []*pollSingleVotesSlideVoteEntry{},
			}
		}

		vote := pollSingleVotesVoteEntry(&poll, &mu, voteMap, optionIndexMap)
		voteEntryGroupsMap[structureLevel.ID].Votes = append(
			voteEntryGroupsMap[structureLevel.ID].Votes,
			&vote,
		)
	}

	structureLevelIDs := make([]int, 0, len(voteEntryGroupsMap))
	for slID := range voteEntryGroupsMap {
		structureLevelIDs = append(structureLevelIDs, slID)
	}

	slices.SortFunc(structureLevelIDs, func(aID, bID int) int {
		nameA := voteEntryGroupsMap[aID].Title
		nameB := voteEntryGroupsMap[bID].Title
		return strings.Compare(nameA, nameB)
	})

	voteEntryGroups := make([]*pollSingleVotesSlideVoteEntryGroup, 0, len(structureLevelIDs))
	for _, slID := range structureLevelIDs {
		voteEntryGroups = append(voteEntryGroups, voteEntryGroupsMap[slID])
	}

	pollMethod := map[string]bool{
		"Yes":     strings.Contains(poll.Pollmethod, "Y"),
		"No":      strings.Contains(poll.Pollmethod, "N"),
		"Abstain": strings.Contains(poll.Pollmethod, "A"),
	}

	slideData.GroupedVotes = voteEntryGroups
	slideData.Options = []*pollSingleVotesSlideOption{}
	slideData.TotalVotesvalid = poll.Votesvalid
	onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
	if !onehundredPercentBase.IsZero() {
		slideData.PercVotesvalid = poll.Votesvalid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
	}

	for _, pollOption := range poll.OptionList {
		option := pollSingleVotesSlideOption{
			TotalYes:     pollOption.Yes,
			TotalNo:      pollOption.No,
			TotalAbstain: pollOption.Abstain,
		}
		if !onehundredPercentBase.IsZero() {
			option.PercYes = option.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercNo = option.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercAbstain = option.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		slideData.Options = append(slideData.Options, &option)
	}

	var globalOption *pollSingleVotesSlideOption
	var globalPollMethod map[string]bool
	if pollOption, ok := poll.GlobalOption.Value(); ok {
		option := pollSingleVotesSlideOption{
			TotalYes:     pollOption.Yes,
			TotalNo:      pollOption.No,
			TotalAbstain: pollOption.Abstain,
		}
		if !onehundredPercentBase.IsZero() {
			option.PercYes = option.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercNo = option.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercAbstain = option.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		globalOption = &option
		globalPollMethod = map[string]bool{
			"Yes":     poll.GlobalYes,
			"No":      poll.GlobalNo,
			"Abstain": poll.GlobalAbstain,
		}
	}

	return map[string]any{
		"_template":        "poll_single_vote",
		"_fullHeight":      true,
		"Data":             slideData,
		"GlobalOption":     globalOption,
		"Title":            poll.Title,
		"LiveVoting":       poll.State == "started" && poll.LiveVotingEnabled,
		"HasResults":       poll.State == "published",
		"Poll":             poll,
		"PollMethod":       pollMethod,
		"GlobalPollMethod": globalPollMethod,
		"NumVotes":         len(voteMap),
		"NumNotVoted":      len(entitledUsers) - len(voteMap),
		"NumEntitledUsers": len(entitledUsers),
		"MaxColumns":       maxColumns,
	}, nil
}

func pollSingleVotesVoteEntry(
	poll *dsmodels.Poll,
	mu *dsmodels.MeetingUser,
	voteMap map[int]string,
	optionIndexMap map[string]int,
) pollSingleVotesSlideVoteEntry {
	user := mu.User
	isPresent := slices.Contains(user.IsPresentInMeetingIDs, poll.MeetingID)
	hasDelegate := false

	if !isPresent && mu.VoteDelegatedTo != nil {
		if delegateMU, ok := mu.VoteDelegatedTo.Value(); ok {
			hasDelegate = slices.Contains(delegateMU.User.IsPresentInMeetingIDs, poll.MeetingID)
		}
	}

	vote := pollSingleVotesSlideVoteEntry{
		FirstName: strings.Trim(user.Title+" "+user.FirstName, " "),
		LastName:  user.LastName,
		Present:   isPresent || hasDelegate,
	}

	if voteVal, ok := voteMap[user.ID]; ok {
		vote.Value = voteVal
		if len(poll.OptionList) > 1 {
			if idx, ok := optionIndexMap[voteVal]; ok {
				vote.Value = strconv.Itoa(idx + 1)
			}
		}
	}

	return vote
}

func mapUsersToVote(poll *dsmodels.Poll) (map[int]string, error) {
	voteMap := map[int]string{}
	if poll.EntitledUsersAtStop != nil {
		globalOption, hasGlobalOption := poll.GlobalOption.Value()
		if hasGlobalOption {
			for _, entry := range globalOption.VoteList {
				if val, ok := entry.UserID.Value(); ok {
					voteMap[val] = entry.Value
				}
			}
		}

		for _, pollOption := range poll.OptionList {
			for _, entry := range pollOption.VoteList {
				if val, ok := entry.UserID.Value(); ok {
					if hasGlobalOption || len(poll.OptionList) > 1 {
						voteMap[val] = strconv.Itoa(pollOption.ID)
					} else {
						voteMap[val] = entry.Value
					}
				}
			}
		}
	} else if poll.LiveVotingEnabled && len(poll.LiveVotes) > 0 {
		pollOption := dsmodels.Option{}
		if len(poll.OptionList) > 0 {
			pollOption = poll.OptionList[0]
		}

		var liveVotes map[int]string
		if err := json.Unmarshal(poll.LiveVotes, &liveVotes); err != nil {
			return nil, fmt.Errorf("parse live votes: %w", err)
		}

		for uid, voteJson := range liveVotes {
			var liveVoteEntry struct {
				RequestUserID int             `json:"request_user_id"`
				VoteUserID    int             `json:"vote_user_id"`
				Value         any             `json:"value"`
				Weight        decimal.Decimal `json:"weight"`
			}
			if err := json.Unmarshal([]byte(voteJson), &liveVoteEntry); err != nil {
				return nil, fmt.Errorf("parse live vote entry: %w", err)
			}

			if voteValue, ok := liveVoteEntry.Value.(map[string]any); ok {
				if len(voteValue) > 1 {
					for optionID, valRaw := range voteValue {
						if val, ok := valRaw.(float64); ok && val == 1 {
							voteMap[uid] = optionID
						}
					}
				} else {
					if val, ok := voteValue[strconv.Itoa(pollOption.ID)]; ok {
						voteMap[uid], _ = val.(string)
					}
				}
			} else if voteValue, ok := liveVoteEntry.Value.(string); ok && voteValue != "" {
				voteMap[uid] = voteValue
			}
		}
	}

	return voteMap, nil
}
