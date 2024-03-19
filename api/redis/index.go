package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func Setcache(key string, value []*model.Item_Autoclik) error {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	expiration := 48 * time.Hour

	err = client.Set(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func Getcache(key string) ([]*model.Item_Autoclik, error) {
	ctx := context.Background()
	client := database.Redis_cache()

	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value []*model.Item_Autoclik
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func Lpush(key string, valuearray []*model.Item_Autoclik) (int64, error) {
	ctx := context.Background()
	client := database.Redis_cache()
	var interfaceSlice []interface{}
	for _, item := range valuearray {
		interfaceSlice = append(interfaceSlice, *item)
	}
	value, err := client.LPush(ctx, key, interfaceSlice...).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}
