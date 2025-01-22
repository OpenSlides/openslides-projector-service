package slide

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"reflect"
	"sort"
	"strconv"
	"strings"

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

	projectionsQ := projectorQ.GetSubquery("current_projection_ids")
	projectionsQ.With("content_object_id", nil)

	var los models.ListOfSpeakers
	losQ := datastore.Collection(req.DB, &models.ListOfSpeakers{}).With("speaker_ids", nil)
	speakersQ := losQ.GetSubquery("speaker_ids")
	meetingUsersQ := speakersQ.With("meeting_user_id", nil)
	meetingUsersQ.With("user_id", nil)

	losSub, err := losQ.SubscribeOne(&los)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe list of speakers: %w", err)
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
				for _, p := range projector.CurrentProjections() {
					if p.ContentObjectID == "" {
						continue
					}

					losId := p.ContentObject().Get("list_of_speakers_id")
					if losId != nil {
						v := reflect.ValueOf(losId)
						if v.Kind() == reflect.Ptr {
							v = v.Elem()
						}

						losQ.SetIds(int(v.Int()))
						if err := losSub.Reload(); err != nil {
							log.Err(err).Msg("Reference projector load failed")
						}
						break
					}
				}
			case <-losSub.Channel:
				if los.ID != 0 {
					content <- getCurrentListOfSpeakersSlideContent(&los, stable)
				}
			}
		}
	}()

	return content, nil
}

func getCurrentListOfSpeakersSlideContent(los *models.ListOfSpeakers, overlay bool) string {
	tmpl, err := template.ParseFiles("templates/slides/current-list-of-speakers.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load current-list-of-speakers template")
		return ""
	}

	type speakerListItem struct {
		Name   string
		Weight int
	}
	waitingSpeakers := []speakerListItem{}
	var currentSpeaker *speakerListItem
	for _, speaker := range los.Speakers() {
		nameParts := []string{}
		if firstName := speaker.MeetingUser().User().FirstName; firstName != nil {
			nameParts = append(nameParts, *firstName)
		}
		if lastName := speaker.MeetingUser().User().LastName; lastName != nil {
			nameParts = append(nameParts, *lastName)
		}

		if len(nameParts) == 0 {
			nameParts = append(nameParts, "User "+strconv.Itoa(speaker.MeetingUser().User().ID))
		}

		weight := 0
		if speaker.Weight != nil {
			weight = *speaker.Weight
		}

		speechState := ""
		if speaker.SpeechState != nil {
			speechState = *speaker.SpeechState
		}

		if (speaker.BeginTime == nil || speechState == "interposed_question") && speaker.EndTime == nil {
			waitingSpeakers = append([]speakerListItem{{
				Name:   strings.Join(nameParts, " "),
				Weight: weight,
			}}, waitingSpeakers...)
		} else if speaker.EndTime == nil || *speaker.EndTime == 0 {
			currentSpeaker = &speakerListItem{
				Name:   strings.Join(nameParts, " "),
				Weight: weight,
			}
		}
	}

	sort.Slice(waitingSpeakers, func(i, j int) bool {
		return waitingSpeakers[i].Weight < waitingSpeakers[j].Weight
	})

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"ListOfSpeakers": los,
		"CurrentSpeaker": currentSpeaker,
		"Speakers":       waitingSpeakers,
		"Overlay":        overlay,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute current-list-of-speakers template")
		return ""
	}

	return content.String()
}
