package redisDB

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "r-uf62qnyehypzinejpepd.redis.rds.aliyuncs.com:6379",
		DB:       0,               //redis默认会创建0-15号DB，这里使用默认的DB
		Username: "zhangqingfeng", //此处不需要用户名
		Password: "Ff85859852",    //没有密码
	})
	//能ping成功才说明连接成功
	if err := Client.Ping(context.Background()).Err(); err != nil {
		fmt.Println("connect to redis failed", err)
	} else {
		fmt.Println("connect to redis")
	}
}
