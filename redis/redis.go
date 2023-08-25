package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Redis *redis.Client

func RedisNewClient(address string, port string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password, // no password set
		DB:       10,       // use default DB
	})

	err := rdb.Set(Ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	val, err := rdb.Get(Ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(Ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	rdb.FlushDB(Ctx)

	return rdb
	// Output: key value
	// key2 does not exist
}
