package yep

import (
	"github.com/hearchco/hearchco/src/search/engines"
)

var Info = engines.Info{
	Domain:         "yep.com",
	Name:           engines.YEP,
	URL:            "https://api.yep.com/fs/2/search?",
	ResultsPerPage: 20,
}

/*
var dompaths = engines.DOMPaths{
	Result:      "div.css-102xgmn-card",
	Link:        "a.css-29ut38-noDecoration",
	Title:       "a.css-29ut38-noDecoration",
	Description: "div.css-1bozosu-snippet",
}
*/

var Support = engines.SupportedSettings{
	Locale:     true,
	SafeSearch: true,
}
