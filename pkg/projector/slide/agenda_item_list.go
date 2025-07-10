package slide

import (
	"context"
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

func AgendaItemListSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

	agendaItemIds, err := req.Fetch.Meeting_AgendaItemIDs(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load agenda item ids %w", err)
	}

	agendaItems, err := req.Fetch.AgendaItem(agendaItemIds...).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load agenda items %w", err)
	}

	agenda, err := recBuildAgendaList(ctx, req.Fetch, agendaItems, 0)
	if err != nil {
		return nil, fmt.Errorf("could process agenda items %w", err)
	}

	return map[string]any{
		"Agenda": agenda,
	}, nil
}

func recBuildAgendaList(ctx context.Context, fetch *dsmodels.Fetch, agendaItems []dsmodels.AgendaItem, currentParent int) ([]agendaListEntry, error) {
	agenda := []agendaListEntry{}
	for _, agendaItem := range agendaItems {
		parentId, _ := agendaItem.ParentID.Value()
		if parentId == currentParent && agendaItem.Type != "internal" && agendaItem.Type != "hidden" {
			titleInfo, err := viewmodels.GetTitleInformationByContentObject(ctx, fetch, agendaItem.ContentObjectID)
			if err != nil {
				return nil, fmt.Errorf("could not get title information: %w", err)
			}

			childEntries, err := recBuildAgendaList(ctx, fetch, agendaItems, agendaItem.ID)
			if err != nil {
				return nil, fmt.Errorf("could not get child entries: %w", err)
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
