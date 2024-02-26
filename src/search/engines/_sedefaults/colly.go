package _sedefaults

import (
	"context"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/hearchco/hearchco/src/config"
	"github.com/hearchco/hearchco/src/search/engines"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func colRequest(ctx context.Context, seName engines.Name) func(r *colly.Request) {
	return func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			if engines.IsTimeoutError(err) {
				log.Trace().
					Err(err).
					Str("engine", seName.String()).
					Msg("_sedefaults.colRequest(): context timeout error")
			} else {
				log.Error().
					Err(err).
					Str("engine", seName.String()).
					Msg("_sedefaults.colRequest(): context error")
			}
			r.Abort()
			return
		}
	}
}

func colError(seName engines.Name) func(r *colly.Response, err error) {
	return func(r *colly.Response, err error) {
		if engines.IsTimeoutError(err) {
			log.Trace().
				// Err(err). // timeout error produces Get "url" error with the query
				Str("engine", seName.String()).
				// Str("url", urll). // can't reliably anonymize it (because it's engine dependent and query isn't passed to this function)
				Msg("_sedefaults.colError(): request timeout error for url")
		} else {
			log.Error().
				Err(err).
				Str("engine", seName.String()).
				// Str("url", urll). // can't reliably anonymize it (because it's engine dependent and query isn't passed to this function)
				Int("statusCode", r.StatusCode).
				Str("response", string(r.Body)). // query can be present, depending on the response from the engine (Google has the query in 3 places)
				Msg("_sedefaults.colError(): request error for url")

			dumpPath := fmt.Sprintf("%v%v_col.log.html", config.LogDumpLocation, seName.String())
			log.Debug().
				Str("engine", seName.String()).
				Str("responsePath", dumpPath).
				Func(func(e *zerolog.Event) {
					bodyWriteErr := os.WriteFile(dumpPath, r.Body, 0644)
					if bodyWriteErr != nil {
						log.Error().
							Err(bodyWriteErr).
							Str("engine", seName.String()).
							Msg("_sedefaults.colError(): error writing html response body to file")
					}
				}).
				Msg("_sedefaults.colError(): html response written")
		}
	}
}
