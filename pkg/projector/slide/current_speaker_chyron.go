package slide

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/database"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
	"github.com/rs/zerolog/log"
)

type currentSpeakerChyronSlideOptions struct {
	ChyronType string `json:"chyron_type"`
	AgendaItem bool   `json:"agenda_item"`
}

func CurrentSpeakerChyronSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	projection := req.Projection

	referenceProjectorId := 0
	refProjectorSub, err := database.Collection(req.DB, &models.Meeting{}).SetFqids(projection.ContentObjectID).SetFields("reference_projector_id").SubscribeField(&referenceProjectorId)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe reference projector id: %w", err)
	}

	var projector models.Projector
	projectorQ := database.Collection(req.DB, &models.Projector{}).With("current_projection_ids", nil)
	projectorSub, err := projectorQ.SubscribeOne(&projector)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe reference projector: %w", err)
	}

	projectionsQ := projectorQ.GetSubquery("current_projection_ids")
	projectionsQ.With("content_object_id", nil)

	var los models.ListOfSpeakers
	losQ := database.Collection(req.DB, &models.ListOfSpeakers{}).With("speaker_ids", nil)
	losQ.With("meeting_id", []string{"list_of_speakers_default_structure_level_time"})
	losContentObjectQ := losQ.With("content_object_id", nil)
	losContentObjectQ.With("agenda_item_id", nil)
	speakersQ := losQ.GetSubquery("speaker_ids")
	sllosQ := speakersQ.With("structure_level_list_of_speakers_id", nil)
	sllosQ.With("structure_level_id", []string{"name"})
	meetingUsersQ := speakersQ.With("meeting_user_id", nil)
	meetingUsersQ.With("user_id", nil)
	meetingUsersQ.With("structure_level_ids", []string{"name"})

	losSub, err := losQ.SubscribeOne(&los)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe list of speakers: %w", err)
	}

	var options currentSpeakerChyronSlideOptions
	if err := json.Unmarshal(projection.Options, &options); err != nil {
		return nil, fmt.Errorf("could not parse slide options: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				refProjectorSub.Unsubscribe()
				projectorSub.Unsubscribe()
				losSub.Unsubscribe()
				close(content)
				return
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
					content <- getSpeakerChyronSlideContent(&los, options)
				} else {
					content <- ""
				}
			}
		}
	}()

	return content, nil
}

func getSpeakerChyronSlideContent(los *models.ListOfSpeakers, options currentSpeakerChyronSlideOptions) string {
	tmpl, err := template.ParseFiles("templates/slides/current-speaker-chyron.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load current-list-of-speakers template")
		return ""
	}

	var currentSpeaker *models.Speaker
	for _, speaker := range los.Speakers() {
		if speaker.IsCurrent() {
			speechState := ""
			if speaker.SpeechState != nil {
				speechState = *speaker.SpeechState
			}

			if speechState == "interposed_question" {
				currentSpeaker = speaker
				break
			} else {
				currentSpeaker = speaker
			}
		}
	}

	speakerName := ""
	structureLevel := ""
	agendaItem := ""
	if currentSpeaker != nil && currentSpeaker.MeetingUser() != nil {
		u := currentSpeaker.MeetingUser().User()
		speakerName = u.ShortName()

		structureLevelDefaultTime := los.Meeting().ListOfSpeakersDefaultStructureLevelTime
		if structureLevelDefaultTime != nil && *structureLevelDefaultTime > 0 {
			sllos := currentSpeaker.StructureLevelListOfSpeakers()
			if sllos != nil {
				structureLevel = sllos.StructureLevel().Name
			}
		} else {
			structureLevelNames := []string{}
			for _, sl := range currentSpeaker.MeetingUser().StructureLevels() {
				structureLevelNames = append(structureLevelNames, sl.Name)
			}

			structureLevel = strings.Join(structureLevelNames, ", ")
		}

		if options.ChyronType == "new" && structureLevel != "" {
			speakerName = fmt.Sprintf("%s, %s", speakerName, structureLevel)
		}
	}

	// TODO: Also include agenda item number and number
	coTitle := los.ContentObject().Get("title")
	if coTitle != nil {
		agendaItem = coTitle.(string)
	}

	slideData := map[string]interface{}{
		"Options":        options,
		"SpeakerName":    speakerName,
		"StructureLevel": structureLevel,
		"AgendaItem":     agendaItem,
	}

	var content bytes.Buffer
	if err := tmpl.Execute(&content, slideData); err != nil {
		log.Error().Err(err).Msg("could not execute current-list-of-speakers template")
		return ""
	}

	return content.String()
}
