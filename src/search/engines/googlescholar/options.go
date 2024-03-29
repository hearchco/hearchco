package googlescholar

import (
	"github.com/hearchco/hearchco/src/search/engines"
)

var Info = engines.Info{
	Domain:         "scholar.google.com",
	Name:           engines.GOOGLESCHOLAR,
	URL:            "https://scholar.google.com/scholar?q=",
	ResultsPerPage: 10,
}

var dompaths = engines.DOMPaths{
	Result:      "div#gs_res_ccl_mid > div.gs_or",
	Link:        "h3 > a",
	Title:       "h3 > a",
	Description: "div.gs_rs",
}

var Support = engines.SupportedSettings{}
