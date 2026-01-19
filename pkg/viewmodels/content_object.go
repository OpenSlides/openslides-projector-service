package viewmodels

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func GetContentObjectField[V any](ctx context.Context, fetch *dsmodels.Fetch, field string, fqid string) (*V, error) {
	dsKey, err := dskey.FromStringf("%s/%s", fqid, field)
	if err != nil {
		return nil, fmt.Errorf("constructing content object dskey: %w", err)
	}

	keys, err := fetch.Get(ctx, dsKey)
	if err != nil {
		return nil, fmt.Errorf("load content object field: %w", err)
	}

	if val, ok := keys[dsKey]; !ok || len(val) == 0 {
		return nil, nil
	}

	var result V
	if err := json.Unmarshal(keys[dsKey], &result); err != nil {
		return nil, fmt.Errorf("parse content object field: %w", err)
	}

	return &result, nil
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
	case "assignemnt", "topic", "motion", "motion_block":
		agendaItemID, err := GetContentObjectField[int](ctx, fetch, "agenda_item_id", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch agenda item id: %w", err)
		}

		if agendaItemID != nil {
			agendaItemNumber, err := fetch.AgendaItem_ItemNumber(*agendaItemID).Value(ctx)
			if err != nil {
				return TitleInformation{}, fmt.Errorf("could not fetch agenda item number: %w", err)
			}

			result.AgendaItemNumber = agendaItemNumber
		}
	}

	// Number
	switch result.Collection {
	case "motion":
		number, err := GetContentObjectField[string](ctx, fetch, "number", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch title: %w", err)
		}

		if number != nil {
			result.Number = *number
		}
	}

	// Title
	switch result.Collection {
	default:
		title, err := GetContentObjectField[string](ctx, fetch, "title", fqid)
		if err != nil {
			return TitleInformation{}, fmt.Errorf("could not fetch title: %w", err)
		}

		if title != nil {
			result.Title = *title
		}
	}

	return result, nil
}
