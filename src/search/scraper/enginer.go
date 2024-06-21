package scraper

import (
	"context"

	"github.com/hearchco/agent/src/config"
	"github.com/hearchco/agent/src/search/engines"
	"github.com/hearchco/agent/src/search/engines/options"
	"github.com/hearchco/agent/src/search/result"
)

// Interface that each search engine must implement to support searching.
type Enginer interface {
	GetName() engines.Name
	GetOrigins() []engines.Name
	Init(context.Context, config.CategoryTimings)
	ReInit(context.Context)
	Search(string, options.Options, chan result.ResultScraped) ([]error, bool)
}