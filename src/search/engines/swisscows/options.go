package swisscows

import (
	"github.com/hearchco/hearchco/src/search/engines"
)

var Info = engines.Info{
	Domain:         "swisscows.com",
	Name:           engines.SWISSCOWS,
	URL:            "https://api.swisscows.com/web/search?",
	ResultsPerPage: 10,
}

var Support = engines.SupportedSettings{
	Locale: true,
}
