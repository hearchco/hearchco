package bing

import (
	"github.com/hearchco/hearchco/src/search/engines"
)

var Info = engines.Info{
	Domain:         "www.bing.com",
	Name:           engines.BING,
	URL:            "https://www.bing.com/search?q=",
	ResultsPerPage: 10,
}

var dompaths = engines.DOMPaths{
	Result:      "ol#b_results > li.b_algo",
	Link:        "h2 > a",
	Title:       "h2 > a",
	Description: "div.b_caption",
}

var Support = engines.SupportedSettings{
	Locale: true,
}
