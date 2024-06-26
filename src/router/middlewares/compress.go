package middlewares

import (
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/klauspost/compress/zstd"
	"github.com/rs/zerolog/log"
)

func compress(lvl int, types ...string) func(next http.Handler) http.Handler {
	// Already has deflate and gzip.
	comp := middleware.NewCompressor(lvl, types...)

	// Add brotli.
	comp.SetEncoder("br", func(w io.Writer, lvl int) io.Writer {
		return brotli.NewWriterOptions(w, brotli.WriterOptions{
			Quality: lvl,
		})
	})

	// Add zstd.
	comp.SetEncoder("zstd", func(w io.Writer, lvl int) io.Writer {
		writer, err := zstd.NewWriter(w, zstd.WithEncoderLevel(zstd.EncoderLevel(lvl)))
		if err != nil {
			log.Panic().Err(err).Msg("Failed to create zstd writer")
		}
		return writer
	})

	return comp.Handler
}
