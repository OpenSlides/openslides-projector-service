package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type meetingMediafileSlideOptions struct {
	Page       int  `json:"page"`
	Fullscreen bool `json:"fullscreen"`
}

func MeetingMediafileSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting mediafile id provided for slide")
	}

	mQ := req.Fetch.MeetingMediafile(*req.ContentObjectID)
	meetingMediafile, err := mQ.Preload(mQ.Mediafile()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load meeting mediafile %w", err)
	}

	options := meetingMediafileSlideOptions{
		Page:       1,
		Fullscreen: true,
	}

	if len(req.Projection.Options) > 0 {
		if err := json.Unmarshal(req.Projection.Options, &options); err != nil {
			return nil, fmt.Errorf("could not parse slide options: %w", err)
		}
	}

	mediafile := meetingMediafile.Mediafile
	fileType := ""
	if mediafile.Mimetype == "application/pdf" {
		fileType = "pdf"
	} else if strings.HasPrefix(mediafile.Mimetype, "image") {
		fileType = "image"
	}

	return map[string]any{
		"MeetingMediafile": meetingMediafile,
		"Mediafile":        meetingMediafile.Mediafile,
		"Options":          options,
		"FileType":         fileType,
	}, nil
}
