package google

import (
	"time"

	"github.com/tminaorg/brzaguza/src/config"
	"github.com/tminaorg/brzaguza/src/structures"
)

var Info structures.SEInfo = structures.SEInfo{
	Domain:         "www.google.com",
	Name:           "Google",
	URL:            "https://www.google.com/search?q=",
	ResultsPerPage: 10,
	Crawlers:       []structures.EngineName{structures.Google},
}

// This should be in SESettings
var timings config.SETimings = config.SETimings{
	Timeout:     10 * time.Second, // the default in colly
	PageTimeout: 5 * time.Second,
	Delay:       100 * time.Millisecond,
	RandomDelay: 50 * time.Millisecond,
	Parallelism: 2, //two requests will be sent to the server, 100 + [0,50) milliseconds apart from the next two
}

var dompaths structures.SEDOMPaths = structures.SEDOMPaths{
	Result:      "div.g",
	Link:        "a",
	Title:       "div > div > div > a > h3",
	Description: "div > div > div > div:first-child > span:first-child",
}

var Support structures.SupportedSettings = structures.SupportedSettings{}