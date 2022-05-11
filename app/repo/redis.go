package repo

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// REDIS null if not initialized
var REDIS *redis.Client

// InitRedis initialize redis connection
func InitRedis() {
	if nil == REDIS {
		redisIndex := 0
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		redisIndex, _ = strconv.Atoi(os.Getenv("REDIS_CACHE"))
		if redisHost != "" {
			REDIS = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       redisIndex,
			})
		}
	}
}
