package viewmodels

import (
	"context"
	"encoding/json"
	"fmt"

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
