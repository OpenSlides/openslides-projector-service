package slide

import (
	"context"
	"fmt"
	"html/template"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
	"github.com/OpenSlides/openslides-projector-service/pkg/viewmodels"
)

func AssignmentSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	aQ := req.Fetch.Assignment(*req.ContentObjectID)
	assignment, err := aQ.Preload(aQ.CandidateList().MeetingUser().StructureLevelList()).Preload(aQ.CandidateList().MeetingUser().User()).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not load assignment id %w", err)
	}

	candidates := []viewmodels.WeightedListEntry{}
	for _, candidate := range assignment.CandidateList {
		var meetingUser *dsmodels.MeetingUser
		if val, isSet := candidate.MeetingUser.Value(); isSet {
			meetingUser = &val
		}

		candidates = append(candidates, viewmodels.WeightedListEntry{
			Name:        req.Locale.Get("Unknown user"),
			MeetingUser: meetingUser,
			Weight:      candidate.Weight,
		})
	}

	viewmodels.CalcWeightedListNames(candidates)

	return map[string]any{
		"Assignment":  assignment,
		"Description": template.HTML(assignment.Description),
		"Candidates":  candidates,
	}, nil
}
