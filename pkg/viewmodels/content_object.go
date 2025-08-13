package viewmodels

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

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

type TitleInformation struct {
	Collection       string
	Title            string
	Number           string
	AgendaItemNumber string
}

func GetTitleInformationByContentObject(ctx context.Context, fetch *dsmodels.Fetch, fqid string) (TitleInformation, error) {
	result := TitleInformation{}

	fqidParts := strings.Split(fqid, "/")
	if len(fqidParts) < 2 {
		return TitleInformation{}, fmt.Errorf("fqid invalid")
	}

	result.Collection = fqidParts[0]

	// AgendaItemNumber
	switch result.Collection {
	case "assignment":
	case "motion":
	case "motion_block":
	case "topic":
		agendaItemID, err := GetContentObjectField[int](ctx, fetch, "agenda_item_id", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch agenda item id: %w", err)
		}

		agendaItemNumber, err := fetch.AgendaItem_ItemNumber(agendaItemID).Value(ctx)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch agenda item number: %w", err)
		}

		result.AgendaItemNumber = agendaItemNumber
	}

	// Number
	switch result.Collection {
	case "motion":
		number, err := GetContentObjectField[string](ctx, fetch, "number", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch title: %w", err)
		}

		result.Number = number
	}

	// Title
	switch result.Collection {
	default:
		title, err := GetContentObjectField[string](ctx, fetch, "title", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch title: %w", err)
		}

		result.Title = title
	}

	return result, nil
}
