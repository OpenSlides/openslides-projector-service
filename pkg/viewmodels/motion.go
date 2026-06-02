package viewmodels

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-go/datastore/dsmodels"
)

func Motion_RecommendationParsed(ctx context.Context, fetch *dsmodels.Fetch, motion *dsmodels.Motion) (string, error) {
	ext := motion.RecommendationExtension
	for _, refMotion := range motion.RecommendationExtensionReferenceIDs {
		title, err := GetTitleInformationByContentObject(ctx, fetch, refMotion)
		if err != nil {
			return "", fmt.Errorf("could not fetch recommendation motion: %w", err)
		}

		name := title.Number
		if name == "" {
			name = title.Title
		}

		ext = strings.ReplaceAll(ext, fmt.Sprintf("[%s]", refMotion), name)
	}

	return ext, nil
}
