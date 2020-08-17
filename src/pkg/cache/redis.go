package cache

import (
	"context"
	"fmt"

	log "goms/pkg/logging"

	"github.com/go-redis/redis/v8"
)

var redisdb *redis.Client
var ctx = context.Background()

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func InitCacheRedis(cfg *RedisConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error("cache redis db", log.ZError(err))
	}
	fmt.Println("client:", client)
	redisdb = client
}

func GetString(key string) []byte {
	res, err := redisdb.Get(ctx, key).Bytes()
	if err != nil {
		log.Error("cache get message", log.ZError(err))
	}
	return res
}

func SetSting(key, value string) string {
	cmd, err := redisdb.Set(ctx, key, value, 0).Result()
	fmt.Println("cmd:", cmd, err)
	if err != nil {
		log.Error("cache set message", log.ZError(err))
	}
	return cmd
}

// 删除redis key内容
func Del(keys []string) int64 {
	cmd, err := redisdb.Del(ctx, keys...).Result()
	if err != nil {
		log.Error("cache del key", log.ZError(err))
	}
	fmt.Println("del:", cmd, err)
	return cmd
}
