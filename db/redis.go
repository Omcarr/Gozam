package db

import (
	"context"
	"encoding/json"
	"fmt"
	"gozam/models"
	"os"
	"strconv"

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

func StoreFingerprints(ctx context.Context, client *redis.Client, songID string, fingerprints map[uint32]models.Couple) error {
	// Convert map[string]uint32 to map[string]string (HSET requires string values)
	data := make(map[string]string, len(fingerprints))
	for k, v := range fingerprints {
		encoded, err := json.Marshal(v)
		if err != nil {
			continue
		}
		data[fmt.Sprint(k)] = string(encoded)
	}

	// Store as Redis hash
	key := "fingerprints:" + songID
	return client.HSet(ctx, key, data).Err()
}

func GetFingerprints(ctx context.Context, client *redis.Client, songID string) (map[uint32]models.Couple, error) {
	key := "fingerprints:" + fmt.Sprint(songID)
	raw, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	result := make(map[uint32]models.Couple)
	for k, v := range raw {
		var couple models.Couple
		if err := json.Unmarshal([]byte(v), &couple); err != nil {
			continue
		}
		id, _ := strconv.ParseUint(k, 10, 32)
		result[uint32(id)] = couple
	}
	return result, nil
}
