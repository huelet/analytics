package src

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectDB() (*redis.Client, error) {
	dbUrl, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	db := redis.NewClient(dbUrl)

	return db, nil
}
