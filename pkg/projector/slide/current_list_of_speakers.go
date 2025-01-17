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

func CurrentListOfSpeakersSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	projection := req.Projection

	var meeting models.Meeting
	q := datastore.Collection(req.DB, &models.Meeting{}).With("speaker_ids", nil).SetFqids(projection.ContentObjectID)
	speakersQ := q.GetSubquery("speaker_ids")
	meetingUsersQ := speakersQ.With("meeting_user_id", nil)
	meetingUsersQ.With("user_id", nil)

	losSub, err := q.SubscribeOne(&meeting)
	if err != nil {
		return nil, fmt.Errorf("CurrentListOfSpeakersSlideHandler: %w", err)
	}

	go func() {
		content <- getCurrentListOfSpeakersSlideContent(&meeting)

		for range <-losSub.Channel {
			content <- getCurrentListOfSpeakersSlideContent(&meeting)
		}
	}()

	return content, nil
}

func getCurrentListOfSpeakersSlideContent(los *models.Meeting) string {
	tmpl, err := template.ParseFiles("templates/slides/current-list-of-speakers.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load current-list-of-speakers template")
		return ""
	}

	speakers := []string{}
	for _, speaker := range los.Speakers() {
		username := speaker.MeetingUser().User().Username
		speakers = append(speakers, username)
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"ListOfSpeakers": los,
		"Speakers":       speakers,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute current-list-of-speakers template")
		return ""
	}

	return content.String()
}
