package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

type agendaListEntry struct {
	Number       string
	TitleInfo    viewmodels.TitleInformation
	Weight       int
	ChildEntries []agendaListEntry
}

type agendaItemListSlideOptions struct {
	OnlyMainItems bool `json:"only_main_items"`
	ShowInternal  bool `json:"-"`
}

func AgendaItemListSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

	options := agendaItemListSlideOptions{
		OnlyMainItems: false,
		ShowInternal:  false,
	}

	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse agenda item list slide options: %w", err)
		}
	}

	req.Fetch.Meeting_AgendaShowInternalItemsOnProjector(req.Projection.MeetingID).Lazy(&options.ShowInternal)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("failed fetching agenda show internal option: %w", err)
	}

	agendaItemIds, err := req.Fetch.Meeting_AgendaItemIDs(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load agenda item ids %w", err)
	}

	agendaItems, err := req.Fetch.AgendaItem(agendaItemIds...).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load agenda items %w", err)
	}

	agenda, err := recBuildAgendaList(ctx, req.Fetch, agendaItems, 0, options)
	if err != nil {
		return nil, fmt.Errorf("could process agenda items %w", err)
	}

	return map[string]any{
		"Agenda": agenda,
	}, nil
}

func recBuildAgendaList(ctx context.Context, fetch *dsmodels.Fetch, agendaItems []dsmodels.AgendaItem, currentParent int, options agendaItemListSlideOptions) ([]agendaListEntry, error) {
	agenda := []agendaListEntry{}
	for _, agendaItem := range agendaItems {
		parentId, _ := agendaItem.ParentID.Value()
		if parentId == currentParent && (options.ShowInternal || agendaItem.Type != "internal") && agendaItem.Type != "hidden" {
			titleInfo, err := viewmodels.GetTitleInformationByContentObject(ctx, fetch, agendaItem.ContentObjectID)
			if err != nil {
				return nil, fmt.Errorf("could not get title information: %w", err)
			}

			var childEntries []agendaListEntry
			if !options.OnlyMainItems {
				childEntries, err = recBuildAgendaList(ctx, fetch, agendaItems, agendaItem.ID, options)
				if err != nil {
					return nil, fmt.Errorf("could not get child entries: %w", err)
				}
			}

			agenda = append(agenda, agendaListEntry{
				Number:       agendaItem.ItemNumber,
				TitleInfo:    titleInfo,
				Weight:       agendaItem.Weight,
				ChildEntries: childEntries,
			})
		}
	}

	slices.SortFunc(agenda, func(a, b agendaListEntry) int {
		return a.Weight - b.Weight
	})

	return agenda, nil
}
