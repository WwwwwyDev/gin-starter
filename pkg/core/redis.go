package core

import (
	"fmt"
	"gin-starter/pkg/global"
	"github.com/go-redis/redis"
	"os"
)

func Redis() *redis.Client{
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("redis connect ping response:"+pong)
	}
	return client
}
