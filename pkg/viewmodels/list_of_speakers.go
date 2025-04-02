package viewmodels

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

func ListOfSpeakers_CurrentSpeaker(ctx context.Context, los *dsfetch.ListOfSpeakers) (*dsfetch.Speaker, error) {
	var currentSpeaker *dsfetch.Speaker
	for _, speakerRef := range los.SpeakerList() {
		speaker, err := speakerRef.Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load speaker: %w", err)
		}

		if Speaker_IsCurrent(&speaker) {
			speechState := speaker.SpeechState

			if speechState == "interposed_question" {
				currentSpeaker = &speaker
				break
			} else {
				currentSpeaker = &speaker
			}
		}
	}

	return currentSpeaker, nil
}
