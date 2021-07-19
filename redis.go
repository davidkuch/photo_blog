package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// docker run -p 6379:6379 redis

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func redisSetSession(name string, uuid string) {
	err := rdb.Set(ctx, uuid, name, 0).Err()
	if err != nil {
		panic(err)
	}
}

func redisGetSession(uuid string) string {
	name, err := rdb.Get(ctx, uuid).Result()
	if err != nil {
		fmt.Println(err)
	}
	return name

}
