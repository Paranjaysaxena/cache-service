package database

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"encoding/json"
)

type redisDatabase struct {
	client *redis.Client
}

type User struct {
	Name     string `json:"name"`
    Class    string `json:"class"`
    RollNum  int64  `json:"roll_num"`
    Metadata []byte `json:"metadata"`
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

func (r *redisDatabase) Set(key string, user User) (User, error) {

	u, _ := json.Marshal(user)

	val, err := r.client.Set(key, u, 0).Result()

	var newUser User

	json.Unmarshal([]byte(val), &newUser)

	return newUser, err
}

func (r *redisDatabase) Get(key string) (User, error) {
	val, err := r.client.Get(key).Result()

	var u User
	json.Unmarshal([]byte(val), &u)

	return u, err
}