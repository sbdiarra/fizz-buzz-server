package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {

	output := zerolog.ConsoleWriter{Out: os.Stderr}

	log.Logger = zerolog.
		New(
			output,
		).With().Str(
		"service", "fizzbuzz",
	).Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
