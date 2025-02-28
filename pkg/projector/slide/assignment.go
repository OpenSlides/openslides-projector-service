package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/rs/zerolog/log"
)

func AssignmentSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string, 1)
	projection := req.Projection

	var assignment models.Assignment
	assignmentSub, err := database.Collection(req.DB, &models.Assignment{}).SetFqids(projection.ContentObjectID).SubscribeOne(&assignment)
	if err != nil {
		return nil, fmt.Errorf("AssignmentSlideHandler: %w", err)
	}

	content <- getAssignmentSlideContent(&assignment)
	go func() {
		for {
			select {
			case <-ctx.Done():
				assignmentSub.Unsubscribe()
				close(content)
				return
			case <-assignmentSub.Channel:
				content <- getAssignmentSlideContent(&assignment)
			}
		}
	}()

	return content, nil
}

func getAssignmentSlideContent(assignment *models.Assignment) string {
	tmpl, err := template.ParseFiles("templates/slides/assignment.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load assignment template")
		return ""
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Assignment": assignment,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute assignment template")
		return ""
	}

	return content.String()
}
