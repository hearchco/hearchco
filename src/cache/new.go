package cache

import (
	"context"
	"fmt"

	"github.com/hearchco/hearchco/src/cache/badger"
	"github.com/hearchco/hearchco/src/cache/nocache"
	"github.com/hearchco/hearchco/src/cache/redis"
	"github.com/hearchco/hearchco/src/config"
	"github.com/rs/zerolog/log"
)

func New(ctx context.Context, fileDbPath string, cacheConf config.Cache) (DB, error) {
	var db DB
	var err error

	switch cacheConf.Type {
	case "badger":
		db, err = badger.New(fileDbPath, cacheConf.Badger)
		if err != nil {
			err = fmt.Errorf("failed creating a badger cache: %w", err)
		}
	case "redis":
		db, err = redis.New(ctx, cacheConf.Redis)
		if err != nil {
			err = fmt.Errorf("failed creating a redis cache: %w", err)
		}
	default:
		db, err = nocache.New()
		if err != nil {
			err = fmt.Errorf("failed creating a nocache: %w", err)
		}
		log.Warn().Msg("Running without caching!")
	}

	return db, err
}
