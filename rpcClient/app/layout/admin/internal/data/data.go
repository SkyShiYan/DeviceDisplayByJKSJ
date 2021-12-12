package data

import (
	"context"
	"rpcClient/app/device/service/ent"
	"rpcClient/app/device/service/internal/conf"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDeviceRepo)

// Data .
type Data struct {
	// wrapped database client
	dbClient *ent.Client
	rClient  *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建DB连接对象
	drv, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	dbClient := ent.NewClient(ent.Driver(drv))
	// 创建表
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		return nil, nil, err
	}

	// 创建Redis连接对象
	rClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		dbClient.Close()
		rClient.Close()
	}
	return &Data{
		dbClient: dbClient,
		rClient:  rClient,
	}, cleanup, nil
}
