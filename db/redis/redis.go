package redis

import (
	"backend/core"
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	log       = core.GetLogger()
	client    *redis.Client
	isReady   = false
	REDIS_URL = os.Getenv("REDIS_URL")
	ctx       = context.Background()
)

func _mem(id string, score int) *redis.Z {
	f := float64(score)
	m := &redis.Z{Member: id, Score: f}
	return m
}
func Start() {
	_opt, _ := redis.ParseURL(REDIS_URL)
	client = redis.NewClient(_opt)

	/*client = redis.NewClient(&redis.Options{
		Addr:     "captain.saas.node.quest:25564",
		Password: "123cache123", // no password set
		DB:       0,             // use default DB
	}) */

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
		// fmt.Println("REDIS ERROR:", err)
	}
	isReady = true
}

func WaitActive() {
	for {
		if !isReady {
			time.Sleep(time.Millisecond * 10)

		}
		if isReady {
			break
		}
	}

}
func GetClient() *redis.Client {
	return client

}

func Get(key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}

		return "", err
	}
	return val, nil
}

func Set(key string, value string) error {
	err := client.Set(ctx, key, value, 0).Err()
	return err
}
