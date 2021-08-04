package database

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

type redisDatabase struct {
	client *redis.Client
}

func createRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	status := client.Ping().Result // makes sure database is connected
	fmt.Println(status)

	return &redisDatabase{client: client}, nil
}

func (r *redisDatabase) Set(key string, value []byte) ([]byte, error) {
	val, err := r.client.Set(key, string(value), 0).Result()
	if err != nil {
		return generateError("set", err)
	}

	return []byte(val), nil
}

func (r *redisDatabase) Get(key string) ([]byte, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return generateError("get", err)
	}

	return []byte(val), nil
}

func generateError(operation string, err error) ([]byte, error) {
	if err == redis.Nil {
		return []byte{}, &OperationError{operation}
	}

	return []byte{}, &DownError{}
}
