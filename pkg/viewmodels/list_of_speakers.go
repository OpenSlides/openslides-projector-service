package viewmodels

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func ListOfSpeakers_CurrentSpeaker(ctx context.Context, los *dsmodels.ListOfSpeakers) (*dsmodels.Speaker, error) {
	var currentSpeaker *dsmodels.Speaker
	for _, speaker := range los.SpeakerList {
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
