package data

import (
	"bff/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBffRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rClient *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建Redis客户端
	rClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cleanup := func() {
		rClient.Close()
		logger.Log(log.LevelInfo, "closing the data resources")
	}
	return &Data{
		rClient: rClient,
	}, cleanup, nil
	// cleanup := func() {
	// 	log.NewHelper(logger).Info("closing the data resources")
	// }
	// return &Data{}, cleanup, nil
}
