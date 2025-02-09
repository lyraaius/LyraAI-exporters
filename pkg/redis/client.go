package redis

import (
	"context"
	"github.com/deagentAI/alphax-exporters/conf"
	goRedis "github.com/redis/go-redis/v9"
)

var AlphaXClient *goRedis.Client

func Init(config *conf.Config) {
	redisConf := config.Redis[0]
	redisClient := goRedis.NewClient(&goRedis.Options{
		Addr:     redisConf.Address,
		Username: redisConf.Username,
		Password: redisConf.Password,
		DB:       redisConf.DB,
		PoolSize: redisConf.PoolSize,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
