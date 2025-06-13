package db

import (
	"context"
	"encoding/json"
	"fmt"
	"gozam/models"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	redisURL := os.Getenv("redisURL")
	if redisURL == "" {
		return nil, fmt.Errorf("redisURL environment variable not set")
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("invalid Redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	return client, nil
}

func StoreFingerprints(ctx context.Context, client *redis.Client, fingerprints map[uint32]models.Couple) error {

	for address, couple := range fingerprints {
		key := fmt.Sprintf("fp:%d", address)
		value, err := json.Marshal(couple)
		if err != nil {
			fmt.Print(err)
		}

		if err := client.Set(ctx, key, value, 0).Err(); err != nil {
			return err // or collect errors if partial failure is OK
		}

	}
	return nil
}

func GetCouples(client *redis.Client, addresses []uint32) (map[uint32][]models.Couple, error) {
	result := make(map[uint32][]models.Couple)

	for _, address := range addresses {
		key := fmt.Sprintf("fp:%d", address)
		ctx := context.Background()

		raw, err := client.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				continue // Key doesn't exist
			}
			return nil, err
		}

		var couple models.Couple
		if err := json.Unmarshal([]byte(raw), &couple); err != nil {
			continue
		}

		result[address] = append(result[address], couple)
	}

	return result, nil
}
