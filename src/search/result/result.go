package result

import (
	"github.com/gocolly/colly/v2"
)

type Image struct {
	Height uint
	Width  uint
}

type ImageResult struct {
	Original     Image
	Thumbnail    Image
	ThumbnailURL string
	Source       string
	SourceURL    string
}

// Everything about some Result, calculated and compiled from multiple search engines
// The URL is the primary key
type Result struct {
	URL           string
	Rank          uint
	Score         float64
	Title         string
	Description   string
	EngineRanks   []RetrievedRank
	TimesReturned uint8
	ImageResult   ImageResult
	Response      *colly.Response
}

func firstN(str string, n int) string {
	v := []rune(str)
	if n >= len(v) {
		return str
	}
	return string(v[:n])
}

func Shorten(results []Result) []Result {
	resultsShort := make([]Result, len(results))
	copy(resultsShort, results)
	for i := range resultsShort {
		if len(resultsShort[i].Description) >= 400 {
			descShort := firstN(resultsShort[i].Description, 397)
			resultsShort[i].Description = descShort + "..."
		}
	}
	return resultsShort
}
