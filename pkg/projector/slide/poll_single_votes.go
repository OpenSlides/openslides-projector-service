package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-go/datastore/dstypes"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSingleVotesSlideVoteEntry struct {
	Value     string
	Present   bool
	Delegated bool
	FirstName string
	LastName  string
}

type pollSingleVotesSlideVoteEntryGroup struct {
	Title string
	Votes []*pollSingleVotesSlideVoteEntry
}

type pollSingleVotesSlideData struct {
	TotalVotesvalid     decimal.Decimal
	PercVotesvalid      decimal.Decimal
	GlobalOption        *pollSingleVotesSlideOption
	GlobalOptionMethods map[string]bool
	Options             []*pollSingleVotesSlideOption
	GroupedVotes        []*pollSingleVotesSlideVoteEntryGroup
}

type pollSingleVotesSlideOption struct {
	ID           int
	Title        string
	Majority     bool
	Weight       int
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
		Preload(pQ.BallotList().PollBallotUser()).
		Preload(pQ.OptionList().ContentObject()).
		Preload(pQ.Config()).
		Preload(pQ.EntitledGroupList().MeetingUserList().VoteDelegatedToList().User().IsPresentInMeetingList()).
		Preload(pQ.EntitledGroupList().MeetingUserList().User().IsPresentInMeetingList()).
		Preload(pQ.EntitledGroupList().MeetingUserList().StructureLevelList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	var maxColumns int
	var nameOrderString dstypes.Meeting_PollProjectionNameOrderFirst
	var usersEnableVoteDelegations bool
	var usersForbidDelegatorToVote bool
	req.Fetch.Meeting_PollProjectionMaxColumns(poll.MeetingID).Lazy(&maxColumns)
	req.Fetch.Meeting_PollProjectionNameOrderFirst(poll.MeetingID).Lazy(&nameOrderString)
	req.Fetch.Meeting_UsersEnableVoteDelegations(poll.MeetingID).Lazy(&usersEnableVoteDelegations)
	req.Fetch.Meeting_UsersForbidDelegatorToVote(poll.MeetingID).Lazy(&usersForbidDelegatorToVote)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not load meeting settings: %w", err)
	}

	var sortByResult bool
	req.Fetch.MeetingPollDefault_SortResultByVotes(poll.MeetingID).Lazy(&sortByResult)

	if nameOrderString == "" {
		nameOrderString = dstypes.Meeting_PollProjectionNameOrderFirstLastName
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

	slices.SortFunc(poll.OptionList, func(a dsmodels.PollOption, b dsmodels.PollOption) int {
		return a.Weight - b.Weight
	})

	optionIndexMap := map[string]int{}
	for idx, option := range poll.OptionList {
		optionIndexMap[strconv.Itoa(option.ID)] = idx
	}

	slideData := pollSingleVotesSlideData{}
	isPublished := poll.Published && poll.State == dstypes.Poll_StateFinished
	if isPublished {
		if err := pollSingleVotesResult(ctx, req, &poll, &slideData); err != nil {
			return nil, fmt.Errorf("calculating poll result: %w", err)
		}

		if len(poll.OptionList) > 1 {
			maxVotes := decimal.Decimal{}
			for _, pollOption := range slideData.Options {
				if maxVotes.LessThan(pollOption.TotalYes) {
					maxVotes = pollOption.TotalYes
				}
			}

			winner := -1
			for oIdx, option := range slideData.Options {
				if option.TotalYes.Equal(maxVotes) {
					// If >1 winners found reset and stop
					if winner != -1 {
						slideData.Options[winner].Majority = false
						idx := strconv.Itoa(slideData.Options[winner].ID)
						for key, val := range voteMap {
							if val == "yes" {
								voteMap[key] = idx
							}
						}
						break
					}

					winner = oIdx
					option.Majority = true
					idx := strconv.Itoa(option.ID)
					for key, val := range voteMap {
						if val == idx {
							voteMap[key] = "yes"
						}
					}
				}
			}

			if sortByResult {
				slices.SortFunc(slideData.Options, func(a, b *pollSingleVotesSlideOption) int {
					if a.Majority && !b.Majority {
						return -1
					} else if b.Majority && !a.Majority {
						return 1
					}
					return a.Weight - b.Weight
				})
			}
		}
	}

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

		vote := pollSingleVotesVoteEntry(&poll, &mu, voteMap, optionIndexMap, usersEnableVoteDelegations, usersForbidDelegatorToVote)
		voteEntryGroupsMap[structureLevel.ID].Votes = append(
			voteEntryGroupsMap[structureLevel.ID].Votes,
			&vote,
		)
	}

	structureLevelIDs := make([]int, 0, len(voteEntryGroupsMap))
	for slID := range voteEntryGroupsMap {
		structureLevelIDs = append(structureLevelIDs, slID)
	}

	slices.Sort(structureLevelIDs)

	voteEntryGroups := make([]*pollSingleVotesSlideVoteEntryGroup, 0, len(structureLevelIDs))
	for _, slID := range structureLevelIDs {
		voteEntryGroups = append(voteEntryGroups, voteEntryGroupsMap[slID])
	}

	pollMethod := pollConfigYNAOptions(poll.Config)
	slideData.GroupedVotes = voteEntryGroups

	/*
		showValidVotesPercent := poll.OnehundredPercentBase != "disabled" &&
			poll.OnehundredPercentBase != "YN" &&
			(slideData.GlobalOption == nil || poll.OnehundredPercentBase[0] != 'Y')

		displayPercAbstain := poll.OnehundredPercentBase ==
			poll.OnehundredPercentBase == "cast" ||
			poll.OnehundredPercentBase == "entitled" ||
			poll.OnehundredPercentBase == "entitled_present" ||
			poll.OnehundredPercentBase == "valid"
	*/

	return map[string]any{
		"_template":    "poll_single_vote",
		"_fullHeight":  true,
		"Data":         slideData,
		"GlobalOption": slideData.GlobalOption,
		// "DisplayPercAbstain":    displayPercAbstain,
		// "GlobalOptionInBase":    poll.OnehundredPercentBase[0] != 'Y' && poll.OnehundredPercentBase != "disabled",
		// "ShowValidVotesPercent": showValidVotesPercent,
		"Title":            poll.Title,
		"LiveVoting":       poll.State == "started" && poll.LiveVotingEnabled,
		"HasResults":       isPublished,
		"HasMultiOptions":  len(poll.OptionList) > 1,
		"Poll":             poll,
		"PollMethod":       pollMethod,
		"GlobalPollMethod": slideData.GlobalOptionMethods,
		"SingleOption":     len(poll.OptionList) == 1,
		"NumVotes":         len(voteMap),
		"NumNotVoted":      len(entitledUsers) - len(voteMap),
		"NumEntitledUsers": len(entitledUsers),
		"MaxColumns":       maxColumns,
	}, nil
}

func pollConfigYNAOptions(config dsmodels.PollConfigUnion) map[string]bool {
	options := map[string]bool{
		"Yes":     false,
		"No":      false,
		"Abstain": false,
	}

	switch config := config.(type) {
	case *dsmodels.PollConfigApproval:
		options["Yes"] = true
		options["No"] = true
		options["Abstain"] = config.AllowAbstain
	case *dsmodels.PollConfigRatingApproval:
		options["Yes"] = true
		options["No"] = true
		options["Abstain"] = config.AllowAbstain
	case *dsmodels.PollConfigRatingScore:
		options["Yes"] = true
	case *dsmodels.PollConfigSelection:
		options["Yes"] = !config.StrikeOut
		options["No"] = config.StrikeOut
	}

	return options
}

func pollSingleVotesVoteEntry(
	poll *dsmodels.Poll,
	mu *dsmodels.MeetingUser,
	voteMap map[int]string,
	optionIndexMap map[string]int,
	usersEnableVoteDelegations bool,
	usersForbidDelegatorToVote bool,
) pollSingleVotesSlideVoteEntry {
	user := mu.User
	isPresent := slices.Contains(user.IsPresentInMeetingIDs, poll.MeetingID)

	hasDelegate := false
	delegatePresent := false

	if len(mu.VoteDelegatedToIDs) > 0 {
		for _, delegateMU := range mu.VoteDelegatedToList {
			delegatePresent = slices.Contains(delegateMU.User.IsPresentInMeetingIDs, poll.MeetingID)
			if delegatePresent {
				hasDelegate = true
			}
		}
	}

	showDelegationIcon := false
	if usersEnableVoteDelegations && hasDelegate {
		if usersForbidDelegatorToVote {
			showDelegationIcon = true
		} else {
			showDelegationIcon = !isPresent
		}
	}

	vote := pollSingleVotesSlideVoteEntry{
		FirstName: strings.Trim(user.Title+" "+user.FirstName, " "),
		LastName:  user.LastName,
		Present:   isPresent || hasDelegate,
		Delegated: showDelegationIcon,
	}

	if voteVal, ok := voteMap[mu.ID]; ok {
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
	for _, ballot := range poll.BallotList {
		if user, ok := ballot.PollBallotUser.Value(); ok {
			if muID, ok := user.RepresentedMeetingUserID.Value(); ok {
				if len(ballot.Value) > 2 {
					voteMap[muID] = ballot.Value[1 : len(ballot.Value)-1]
				}
			}
		}
	}

	return voteMap, nil
}

func pollSingleVotesResult(
	ctx context.Context,
	req *projectionRequest,
	poll *dsmodels.Poll,
	data *pollSingleVotesSlideData,
) error {
	data.Options = []*pollSingleVotesSlideOption{}

	switch poll.Config.(type) {
	case *dsmodels.PollConfigApproval:
		var result viewmodels.PollResultApproval
		if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
			return fmt.Errorf("parse approval poll result %w", err)
		}

		data.TotalVotesvalid = decimal.NewFromInt(result.VotesValid())

		onehundredPercentBase := result.OneHundredPercentBase(poll.Config.(*dsmodels.PollConfigApproval))
		if !onehundredPercentBase.IsZero() {
			data.PercVotesvalid = data.TotalVotesvalid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		option := &pollSingleVotesSlideOption{
			TotalYes:     result.Yes,
			TotalNo:      result.No,
			TotalAbstain: result.Abstain,
			Majority:     false,
		}

		if !onehundredPercentBase.IsZero() {
			option.PercYes = option.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercNo = option.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			option.PercAbstain = option.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		data.Options = append(data.Options, option)
	case *dsmodels.PollConfigSelection:
		config := poll.Config.(*dsmodels.PollConfigSelection)
		var result viewmodels.PollResultSelection
		if err := json.Unmarshal([]byte(poll.Result), &result); err != nil {
			return fmt.Errorf("parse approval poll result %w", err)
		}

		data.TotalVotesvalid = decimal.NewFromInt(result.VotesValid())

		onehundredPercentBase := result.OneHundredPercentBase(config)
		if !onehundredPercentBase.IsZero() {
			data.PercVotesvalid = data.TotalVotesvalid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		}

		for _, option := range poll.OptionList {
			optionLabel, err := viewmodels.Option_OptionLabel(ctx, req.Fetch, req.Locale, &option)
			if err != nil {
				return fmt.Errorf("parsing option label: %w", err)
			}

			option := &pollSingleVotesSlideOption{
				ID:       option.ID,
				Title:    optionLabel,
				TotalYes: result.Options[strconv.Itoa(option.ID)],
				Weight:   option.Weight,
				Majority: false,
			}

			if !onehundredPercentBase.IsZero() {
				option.PercYes = option.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
				option.PercNo = option.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
				option.PercAbstain = option.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			}

			data.Options = append(data.Options, option)
		}

		if config.MinOptionsAmount == 0 || config.AllowNota {
			option := pollSingleVotesSlideOption{
				TotalAbstain: result.Abstain,
			}

			if config.AllowNota && config.StrikeOut {
				option.TotalYes = result.Nota
			} else if config.AllowNota && !config.StrikeOut {
				option.TotalNo = result.Nota
			}

			if !onehundredPercentBase.IsZero() && config.OnehundredPercentBase[:3] != "yes" {
				option.PercYes = option.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
				option.PercNo = option.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
				option.PercAbstain = option.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
			}

			data.GlobalOption = &option
			data.GlobalOptionMethods = map[string]bool{
				"Yes":     config.AllowNota && config.StrikeOut,
				"No":      config.AllowNota && !config.StrikeOut,
				"Abstain": config.MinOptionsAmount == 0,
			}
		}
	}

	return nil
}
