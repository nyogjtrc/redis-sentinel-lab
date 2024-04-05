package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	cli := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",
		SentinelAddrs: []string{
			"sentinel-1:26379",
			"sentinel-2:26379",
			"sentinel-3:26379",
		},
	})

	for {
		result, err := cli.Ping(ctx).Result()
		fmt.Println("redis ping", result, err)

		info := cli.ClientInfo(ctx).Val()
		if info != nil {
			fmt.Println("client info addr", info.Addr)
			fmt.Println("client info laddr", info.LAddr)
		}

		time.Sleep(2 * time.Second)
	}
}
