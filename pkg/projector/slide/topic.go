package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/rs/zerolog/log"
)

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	projection := req.Projection

	var topic models.Topic
	topicSub, err := datastore.Collection(req.DB, &models.Topic{}).With("agenda_item_id", nil).SetFqids(projection.ContentObjectID).SubscribeOne(&topic)
	if err != nil {
		return nil, fmt.Errorf("TopicSlideHandler: %w", err)
	}

	go func() {
		content <- getTopicSlideContent(&topic)

		for range <-topicSub.Channel {
			content <- getTopicSlideContent(&topic)
		}
	}()

	return content, nil
}

func getTopicSlideContent(topic *models.Topic) string {
	tmpl, err := template.ParseFiles("templates/slides/topic.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load topic template")
		return ""
	}

	text := ""
	if topic.Text != nil {
		text = *topic.Text
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"AgendaItem": topic.AgendaItem(),
		"Topic":      topic,
		"Text":       template.HTML(text),
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute topic template")
		return ""
	}

	return content.String()
}
