package slide

import (
	"context"
	"fmt"
	"html/template"
)

func HomeSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

	welcomeTitle := ""
	welcomeText := ""
	req.Fetch.Meeting_WelcomeTitle(*req.ContentObjectID).Lazy(&welcomeTitle)
	req.Fetch.Meeting_WelcomeText(*req.ContentObjectID).Lazy(&welcomeText)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not fetch wlan data")
	}

	return map[string]any{
		"Title": welcomeTitle,
		"Text":  template.HTML(welcomeText),
	}, nil
}
