package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
	"github.com/shopspring/decimal"
)

type pollSlideOptions struct {
	SingleVotes bool `json:"single_votes"`
}

func PollSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	options := pollSlideOptions{
		SingleVotes: false,
	}
	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	if options.SingleVotes {
		return pollSingleVotesSlideHandler(ctx, req)
	}

	poll, err := req.Fetch.Poll(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	title, err := viewmodels.GetContentObjectField[string](ctx, req.Fetch, "title", poll.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load poll content object title %w", err)
	}

	return map[string]any{
		"Title": title,
		"Poll":  poll,
	}, nil
}

type pollSingleVotesSlideVoteEntry struct {
	Value     string
	FirstName string
	LastName  string
}

type pollSingleVotesSlideVoteEntryGroup struct {
	TotalYes     decimal.Decimal
	TotalNo      decimal.Decimal
	TotalAbstain decimal.Decimal
	Votes        map[int]*pollSingleVotesSlideVoteEntry
}

type pollSingleVotesSlideData struct {
	TotalYes     decimal.Decimal
	TotalNo      decimal.Decimal
	TotalAbstain decimal.Decimal
	PercYes      decimal.Decimal
	PercNo       decimal.Decimal
	PercAbstain  decimal.Decimal
	Votesvalid   decimal.Decimal
	GroupedVotes pollSingleVotesSlideVoteEntryGroup
}

func pollSingleVotesSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	pQ := req.Fetch.Poll()
	poll, err := req.Fetch.Poll(*req.ContentObjectID).Preload(pQ.OptionList().VoteList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load poll id %w", err)
	}

	pollOption := dsmodels.Option{}
	if len(poll.OptionList) > 0 {
		pollOption = poll.OptionList[0]
	}

	var entitledUsersAtStop []struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(poll.EntitledUsersAtStop, &entitledUsersAtStop); err != nil {
		return nil, fmt.Errorf("parse los id: %w", err)
	}

	numEntitledUsers := len(entitledUsersAtStop)
	pollMethod := map[string]bool{
		"Yes":     strings.Contains(poll.Pollmethod, "Y"),
		"No":      strings.Contains(poll.Pollmethod, "N"),
		"Abstain": strings.Contains(poll.Pollmethod, "A"),
	}

	type voteEntry struct {
		pollSingleVotesSlideVoteEntry
	}

	voteEntries := map[int]*voteEntry{}
	for _, entry := range entitledUsersAtStop {
		user, err := req.Fetch.User(entry.UserID).First(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load entitled user: %w", err)
		}

		vote := pollSingleVotesSlideVoteEntry{
			FirstName: strings.Trim(user.Title+" "+user.FirstName, " "),
			LastName:  user.LastName,
			Value:     "",
		}
		voteEntries[entry.UserID] = &voteEntry{
			pollSingleVotesSlideVoteEntry: vote,
		}
	}

	for _, entry := range pollOption.VoteList {
		if val, ok := entry.UserID.Value(); ok {
			voteEntries[val].Value = entry.Value
		}
	}

	slideData := pollSingleVotesSlideData{}
	slideData.TotalYes, _ = pollOption.Yes.Value()
	slideData.TotalNo, _ = pollOption.No.Value()
	slideData.TotalAbstain, _ = pollOption.Abstain.Value()

	return map[string]any{
		"_template":        "poll_single_vote",
		"_fullHeight":      true,
		"Data":             slideData,
		"Title":            poll.Title,
		"Poll":             poll,
		"PollMethod":       pollMethod,
		"PollOption":       pollOption,
		"NumEntitledUsers": numEntitledUsers,
		"Votes":            voteEntries,
	}, nil
}
