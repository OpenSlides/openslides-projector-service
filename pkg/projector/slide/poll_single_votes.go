package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
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
	TotalYes        decimal.Decimal
	TotalNo         decimal.Decimal
	TotalAbstain    decimal.Decimal
	TotalVotesvalid decimal.Decimal
	PercYes         decimal.Decimal
	PercNo          decimal.Decimal
	PercAbstain     decimal.Decimal
	PercVotesvalid  decimal.Decimal
	GroupedVotes    []*pollSingleVotesSlideVoteEntryGroup
}

func pollSingleVotesSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pQ := req.Fetch.Poll()
	poll, err := req.Fetch.Poll(*req.ContentObjectID).
		Preload(pQ.OptionList().VoteList()).
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

	pollOption := dsmodels.Option{}
	if len(poll.OptionList) > 0 {
		pollOption = poll.OptionList[0]
	}

	voteMap := map[int]string{}
	if poll.EntitledUsersAtStop != nil {
		for _, entry := range pollOption.VoteList {
			if val, ok := entry.UserID.Value(); ok {
				voteMap[val] = entry.Value
			}
		}
	} else if poll.LiveVotingEnabled {
		var liveVotes map[int]string
		if err := json.Unmarshal(poll.LiveVotes, &liveVotes); err != nil {
			return nil, fmt.Errorf("parse los id: %w", err)
		}

		for uid, voteJson := range liveVotes {
			var liveVoteEntry struct {
				RequestUserID int             `json:"request_user_id"`
				VoteUserID    int             `json:"vote_user_id"`
				Value         map[int]string  `json:"value"`
				Weight        decimal.Decimal `json:"weight"`
			}
			if err := json.Unmarshal([]byte(voteJson), &liveVoteEntry); err != nil {
				return nil, fmt.Errorf("parse los id: %w", err)
			}

			if val, ok := liveVoteEntry.Value[pollOption.ID]; ok {
				voteMap[uid] = val
			}
		}
	}

	meetingUserMap := map[int]dsmodels.MeetingUser{}
	for _, group := range poll.EntitledGroupList {
		for _, mu := range group.MeetingUserList {
			meetingUserMap[mu.UserID] = mu
		}
	}

	slideData := pollSingleVotesSlideData{}
	voteEntryGroupsMap := map[int]*pollSingleVotesSlideVoteEntryGroup{}
	entitledUsers := viewmodels.Poll_EntitledUserIDsSorted(poll, nameOrderString)
	for _, userID := range entitledUsers {
		mu, exists := meetingUserMap[userID]
		if !exists {
			continue
		}

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
			Value:     voteMap[user.ID],
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

	pollMethod := map[string]bool{
		"Yes":     strings.Contains(poll.Pollmethod, "Y"),
		"No":      strings.Contains(poll.Pollmethod, "N"),
		"Abstain": strings.Contains(poll.Pollmethod, "A"),
	}

	slideData.GroupedVotes = voteEntryGroups
	slideData.TotalYes = pollOption.Yes
	slideData.TotalNo = pollOption.No
	slideData.TotalAbstain = pollOption.Abstain
	slideData.TotalVotesvalid = poll.Votesvalid
	onehundredPercentBase := viewmodels.Poll_OneHundredPercentBase(poll, nil)
	if !onehundredPercentBase.IsZero() {
		slideData.PercYes = slideData.TotalYes.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		slideData.PercNo = slideData.TotalNo.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		slideData.PercAbstain = slideData.TotalAbstain.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
		slideData.PercVotesvalid = slideData.TotalVotesvalid.DivRound(onehundredPercentBase, 5).Mul(decimal.NewFromInt(100))
	}

	return map[string]any{
		"_template":        "poll_single_vote",
		"_fullHeight":      true,
		"Data":             slideData,
		"Title":            poll.Title,
		"LiveVoting":       poll.State == "started" && poll.LiveVotingEnabled,
		"Poll":             poll,
		"PollMethod":       pollMethod,
		"PollOption":       pollOption,
		"NumVotes":         len(voteMap),
		"NumNotVoted":      len(entitledUsers) - len(voteMap),
		"NumEntitledUsers": len(entitledUsers),
		"MaxColumns":       maxColumns,
	}, nil
}
