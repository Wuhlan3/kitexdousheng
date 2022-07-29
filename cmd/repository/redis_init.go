package repository

import (
	"context"
	"kitexdousheng/pkg/constants"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

// 初始化连接
func RedisInit() (err error) {
	constants.REDIS = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "Ling18", // no password set
		DB:       0,        // use default DB
		PoolSize: 100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = constants.REDIS.Ping(ctx).Result()
	return err
}
