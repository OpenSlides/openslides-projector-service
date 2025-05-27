package slide

import (
	"context"
	"fmt"
)

func AssignmentSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	assignment, err := req.Fetch.Assignment(*req.ContentObjectID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load assignment id %w", err)
	}

	return map[string]any{
		"Assignment": assignment,
	}, nil
}
