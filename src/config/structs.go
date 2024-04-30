package config

import (
	"time"

	"github.com/hearchco/hearchco/src/search/category"
	"github.com/hearchco/hearchco/src/search/engines"
)

type EngineRanking struct {
	Mul   float64 `koanf:"mul"`
	Const float64 `koanf:"const"`
}

type Ranking struct {
	REXP    float64                  `koanf:"rexp"`
	A       float64                  `koanf:"a"`
	B       float64                  `koanf:"b"`
	C       float64                  `koanf:"c"`
	D       float64                  `koanf:"d"`
	TRA     float64                  `koanf:"tra"`
	TRB     float64                  `koanf:"trb"`
	TRC     float64                  `koanf:"trc"`
	TRD     float64                  `koanf:"trd"`
	Engines map[string]EngineRanking `koanf:"engines"`
}

type Settings struct {
	RequestedResultsPerPage int      `koanf:"requestedresults"`
	Shortcut                string   `koanf:"shortcut"`
	Proxies                 []string `koanf:"proxies"`
}

// ReaderTTL is format in which the config is read from the config file
// in <number><unit> format
// example: 1s, 1m, 1h, 1d, 1w, 1M, 1y
// if unit is not specified, it is assumed to be milliseconds
type ReaderTTL struct {
	// how long to store the results in cache
	// setting this to 0 caches the results forever
	// to disable caching set conf.Cache.Type to "none"
	Time string `koanf:"time"`
	// if the remaining TTL when retrieving from cache is less than this, update the cache entry and reset the TTL
	// setting this to 0 disables this feature
	// setting this to the same value (or higher) as Results will update the cache entry every time
	RefreshTime string `koanf:"refreshtime"`
}
type TTL struct {
	Time        time.Duration
	RefreshTime time.Duration
}

// ReaderCache is format in which the config is read from the config file
type ReaderCache struct {
	// can be "none", "badger" or "redis"
	Type string `koanf:"type"`
	// has no effect if Type is "none"
	TTL ReaderTTL `koanf:"ttl"`
}
type Cache struct {
	Type string
	TTL  TTL
}

type ReaderProxyTimeouts struct {
	Dial         string `koanf:"dial"`
	KeepAlive    string `koanf:"keepalive"`
	TLSHandshake string `koanf:"tlshandshake"`
}
type ProxyTimeouts struct {
	Dial         time.Duration
	KeepAlive    time.Duration
	TLSHandshake time.Duration
}

type ReaderProxy struct {
	Salt     string              `koanf:"salt"`
	Timeouts ReaderProxyTimeouts `koanf:"timeouts"`
}
type Proxy struct {
	Salt     string
	Timeouts ProxyTimeouts
}

// ReaderServer is format in which the config is read from the config file
type ReaderServer struct {
	// environment in which the server is running (normal or lambda)
	Environment string `koanf:"environment"`
	// port on which the API server listens
	Port int `koanf:"port"`
	// urls used for CORS, comma separated (wildcards allowed) and converted into slice
	FrontendUrls string `koanf:"frontendurls"`
	// cache settings
	Cache ReaderCache `koanf:"cache"`
	// salt used for image proxy
	Proxy ReaderProxy `koanf:"proxy"`
}
type Server struct {
	Environment  string
	Port         int
	FrontendUrls []string
	Cache        Cache
	Proxy        Proxy
}

// ReaderEngine is format in which the config is read from the config file
type ReaderEngine struct {
	Enabled bool `koanf:"enabled"`
}

// ReaderTimings is format in which the config is read from the config file
// in miliseconds
type ReaderTimings struct {
	PreferredTimeout        uint `koanf:"preferredtimeout"`
	PreferredTimeoutResults int  `koanf:"preferredtimeoutresults"`
	AdditionalTimeout       uint `koanf:"additionaltimeout"`
	HardTimeout             uint `koanf:"hardtimeout"`
	Timeout                 uint `koanf:"timeout"`
	PageTimeout             uint `koanf:"pagetimeout"`
	Delay                   uint `koanf:"delay"`
	RandomDelay             uint `koanf:"randomdelay"`
	Parallelism             int  `koanf:"parallelism"`
}

// Delegates Timeout, PageTimeout to colly.Collector.SetRequestTimeout(); Note: See https://github.com/gocolly/colly/issues/644
// Delegates Delay, RandomDelay, Parallelism to colly.Collector.Limit()
type Timings struct {
	// preferred timeout if enough results are found
	PreferredTimeout time.Duration
	// number of results which if not met will trigger the additional timeout
	PreferredTimeoutResults int
	// additional timeout if not enough results are found (delay after which the number of results is checked)
	AdditionalTimeout time.Duration
	// hard timeout after which the search is forcefully stopped
	HardTimeout time.Duration
	// colly settings
	Timeout     time.Duration
	PageTimeout time.Duration
	Delay       time.Duration
	RandomDelay time.Duration
	Parallelism int
}

// ReaderCategory is format in which the config is read from the config file
type ReaderCategory struct {
	REngines map[string]ReaderEngine `koanf:"engines"`
	Ranking  Ranking                 `koanf:"ranking"`
	RTimings ReaderTimings           `koanf:"timings"`
}
type Category struct {
	Engines []engines.Name
	Ranking Ranking
	Timings Timings
}

// ReaderConfig is format in which the config is read from the config file
type ReaderConfig struct {
	Server      ReaderServer                     `koanf:"server"`
	RCategories map[category.Name]ReaderCategory `koanf:"categories"`
	Settings    map[string]Settings              `koanf:"settings"`
}
type Config struct {
	Server     Server
	Categories map[category.Name]Category
	Settings   map[engines.Name]Settings
}
