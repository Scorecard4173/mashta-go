package mashta

import (
	"github.com/rs/zerolog/log"
)

func MessageHandler(applicationContext ApplicationContext, handlerFunc HandlerFunc, retrier MessageRetrier) func(event MessageEvent) {
	return func(event MessageEvent) {
		status := handlerFunc(event)
		switch status {
		case ProcessingSuccess:
			log.Info().Msg("successfully processed message")
		case SkipMessage:
			log.Info().Msg("skipping message")
		case RetryMessage:
			log.Info().Msgf("retrying message")
			if retryErr := retrier.Retry(applicationContext, event); retryErr != nil {
				log.Error().Err(retryErr).Msg("error retrying message")
			}
		}
	}
}
