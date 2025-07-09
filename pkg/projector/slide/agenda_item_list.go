package slide

import (
	"context"
	"fmt"
	"slices"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

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

	// TODO: Fix sorting
	slices.SortFunc(agendaItems, func(a, b dsmodels.AgendaItem) int {
		parentIdA, _ := a.ParentID.Value()
		parentIdB, _ := b.ParentID.Value()
		if parentIdA == parentIdB {
			if a.Weight == b.Weight {
				return a.ID - b.ID
			}
			return a.Weight - b.Weight
		} else if a.ID == parentIdB {
			return -1
		} else if b.ID == parentIdA {
			return 1
		}

		println(parentIdA, parentIdB, a.ID, b.ID)

		return -1
	})

	type agendaListEntry struct {
		Number string
		Title  string
		Type   string
		Level  int
		Weight int
	}
	agenda := []agendaListEntry{}
	for _, agendaItem := range agendaItems {
		agenda = append(agenda, agendaListEntry{
			Number: agendaItem.ItemNumber,
			Level:  agendaItem.Level,
			Weight: agendaItem.Weight,
		})
	}

	return map[string]any{
		"Agenda": agenda,
	}, nil
}
