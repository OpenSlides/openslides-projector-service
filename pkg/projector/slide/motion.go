package slide

/*
func MotionSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string, 1)
	projection := req.Projection

	var motion models.Motion
	motionSub, err := database.Collection(req.DB, &models.Motion{}).SetFqids(projection.ContentObjectID).SubscribeOne(&motion)
	if err != nil {
		return nil, fmt.Errorf("MotionSlideHandler: %w", err)
	}

	content <- getMotionSlideContent(&motion)
	go func() {
		for {
			select {
			case <-ctx.Done():
				motionSub.Unsubscribe()
				close(content)
				return
			case <-motionSub.Channel:
				content <- getMotionSlideContent(&motion)
			}
		}
	}()

	return content, nil
}

func getMotionSlideContent(motion *models.Motion) string {
	tmpl, err := template.ParseFiles("templates/slides/motion.html")
	if err != nil {
		log.Error().Err(err).Msg("could not load motion template")
		return ""
	}

	var content bytes.Buffer
	err = tmpl.Execute(&content, map[string]interface{}{
		"Motion": motion,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not execute motion template")
		return ""
	}

	return content.String()
}
*/
