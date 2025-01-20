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

	referenceProjectorId := 0
	refProjectorSub, err := datastore.Collection(req.DB, &models.Meeting{}).SetFqids(projection.ContentObjectID).SetFields("reference_projector_id").SubscribeField(&referenceProjectorId)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe reference projector id: %w", err)
	}

	var projector models.Projector
	projectorQ := datastore.Collection(req.DB, &models.Projector{}).With("current_projection_ids", nil)
	projectorSub, err := projectorQ.SubscribeOne(&projector)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe reference projector: %w", err)
	}

	stable := false
	if projection.Stable != nil {
		stable = *projection.Stable
	}

	go func() {
		for {
			select {
			case <-refProjectorSub.Channel:
				if referenceProjectorId > 0 {
					projectorQ.SetIds(referenceProjectorId)
					if err := projectorSub.Reload(); err != nil {
						log.Err(err).Msg("Reference projector load failed")
					}
				}
			case <-projectorSub.Channel:
				for _, projection := range projector.CurrentProjections() {
					if projection.ContentObjectID == "" {
						continue;
					}

					println(projection.ContentObjectID, stable)
				}
			}
		}
	}()

	/*
	var projector models.Projector
	q := datastore.Collection(req.DB, &models.Projector{}).With("current_projection_ids", nil).SetIds(referenceProjectorId)
	// speakersQ := q.GetSubquery("speaker_ids")
	// meetingUsersQ := speakersQ.With("meeting_user_id", nil)
	// meetingUsersQ.With("user_id", nil)

	losSub, err := q.SubscribeOne(&meeting)
	if err != nil {
		return nil, fmt.Errorf("CurrentListOfSpeakersSlideHandler: %w", err)
	}

	go func() {
		content <- getCurrentListOfSpeakersSlideContent(&meeting, stable)

		for range <-losSub.Channel {
			content <- getCurrentListOfSpeakersSlideContent(&meeting, stable)
		}
	}()
	*/

	return content, nil
}

func getCurrentListOfSpeakersSlideContent(los *models.ListOfSpeakers, overlay bool) string {
	tmpl, err := template.ParseFiles("templates/slides/current-list-of-speakers.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load current-list-of-speakers template")
		return ""
	}

	type speakerListItem struct {
		Number int
		Name string
	}
	speakers := []speakerListItem{}
	for i, speaker := range los.Speakers() {
		username := speaker.MeetingUser().User().Username
		speakers = append(speakers, speakerListItem{ Number: i + 1, Name: username })
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"ListOfSpeakers": los,
		"Speakers":       speakers,
		"Overlay":        overlay,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute current-list-of-speakers template")
		return ""
	}

	return content.String()
}
