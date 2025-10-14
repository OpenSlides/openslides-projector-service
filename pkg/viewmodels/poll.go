package viewmodels

import (
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/rs/zerolog/log"
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
		log.Error().Msg("invalid request for onehundred percent base calculation")
		return decimal.Decimal{}
	}

	switch poll.OnehundredPercentBase {
	case "Y":
		total, _ := poll.Votesvalid.Value()
		if globalOption, isSet := poll.GlobalOption.Value(); isSet {
			abstain, _ := globalOption.Abstain.Value()
			no, _ := globalOption.No.Value()
			total = total.Sub(no).Sub(abstain)
		}
		return total
	case "YN":
		yes, _ := option.Yes.Value()
		no, _ := option.No.Value()
		return yes.Add(no)
	case "YNA":
		yes, _ := option.Yes.Value()
		no, _ := option.No.Value()
		abstain, _ := option.Abstain.Value()

		return yes.Add(no).Add(abstain)
	case "valid":
		valid, _ := poll.Votesvalid.Value()
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
		cast, _ := poll.Votescast.Value()
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
