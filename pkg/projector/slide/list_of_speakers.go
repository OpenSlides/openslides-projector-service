package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func ListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no list of speakers id provided for slide")
	}

	t := req.Fetch.ListOfSpeakers(*req.ContentObjectID)
	los, err := t.Preload(t.SpeakerList()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	titleInfo, err := viewmodels.GetTitleInformationByContentObject(ctx, req.Fetch, los.ContentObjectID)
	if err != nil {
		return nil, fmt.Errorf("could not load los title info %w", err)
	}

	return map[string]any{
		"LoS":          los,
		"ContentTitle": titleInfo,
	}, nil
}
