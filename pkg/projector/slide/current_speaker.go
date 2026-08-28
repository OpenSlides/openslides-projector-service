package slide

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dstypes"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func CurrentSpeakerSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

	referenceProjectorId, err := req.Fetch.Meeting_ReferenceProjectorID(*req.ContentObjectID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load reference projector id %w", err)
	}

	losID, err := viewmodels.Projector_ListOfSpeakersID(ctx, req.Fetch, referenceProjectorId)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers id %w", err)
	}

	if losID == nil {
		return nil, nil
	}

	l := req.Fetch.ListOfSpeakers(*losID)
	los, err := l.Preload(l.SpeakerList().StructureLevelListOfSpeakers().StructureLevel()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load list of speakers %w", err)
	}

	speaker, err := viewmodels.ListOfSpeakers_CurrentSpeaker(ctx, &los)
	if err != nil {
		return nil, fmt.Errorf("could not fetch current speaker %w", err)
	}

	if speaker == nil {
		return nil, nil
	}

	type speakerInfo struct {
		ID                 int
		Name               string
		Color              string
		CountdownTime      float64
		RemainingTime      *float64
		Running            bool
		Intervention       bool
		Answer             bool
		InterposedQuestion bool
		Forspeech          bool
		Counterspeech      bool
		Contribution       bool
		PointOfOrder       bool
	}

	var slideSpeaker speakerInfo
	slideSpeaker.Running = speaker.PauseTime == 0
	slideSpeaker.Intervention = speaker.SpeechState == dstypes.Speaker_SpeechStateIntervention
	slideSpeaker.InterposedQuestion = speaker.SpeechState == dstypes.Speaker_SpeechStateInterposedQuestion
	slideSpeaker.Counterspeech = speaker.SpeechState == dstypes.Speaker_SpeechStateContra
	slideSpeaker.Forspeech = speaker.SpeechState == dstypes.Speaker_SpeechStatePro
	slideSpeaker.Contribution = speaker.SpeechState == dstypes.Speaker_SpeechStateContribution
	slideSpeaker.PointOfOrder = speaker.PointOfOrder
	slideSpeaker.Answer = speaker.Answer

	sllos, hasSLLOS := speaker.StructureLevelListOfSpeakers.Value()
	if hasSLLOS {
		slideSpeaker.ID = sllos.StructureLevelID
		slideSpeaker.Color = sllos.StructureLevel.Color
	}

	if slideSpeaker.InterposedQuestion || speaker.Answer {
		slideSpeaker.CountdownTime = viewmodels.Speaker_CalculateElapsedTime(speaker)
		if hasSLLOS {
			slideSpeaker.Name = sllos.StructureLevel.Name
		}
	} else if slideSpeaker.Intervention {
		defaultInterventionTime, err := req.Fetch.Meeting_ListOfSpeakersInterventionTime(los.MeetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not load intervention time: %w", err)
		}
		slideSpeaker.CountdownTime = viewmodels.Speaker_CalculateInterventionCountdownTime(speaker, defaultInterventionTime)
		slideSpeaker.RemainingTime = new(float64(defaultInterventionTime))
	} else if hasSLLOS {
		slideSpeaker.Name = sllos.StructureLevel.Name
		slideSpeaker.CountdownTime = sllos.RemainingTime + float64(sllos.CurrentStartTime)
		slideSpeaker.RemainingTime = &sllos.RemainingTime
	} else {
		slideSpeaker.CountdownTime = viewmodels.Speaker_CalculateInterventionCountdownTime(speaker, 0)
	}

	return map[string]any{
		"_template":   "current_speaker",
		"SpeakerInfo": slideSpeaker,
	}, nil
}
