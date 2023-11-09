package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/Yeuoly/coshub/internal/static"
	"github.com/Yeuoly/coshub/internal/utils/typ"
	redis "github.com/go-redis/redis/v8"
)

var (
	redis_host                     = "127.0.0.1"
	redis_port                     = 6379
	redis_pass                     = ""
	redis_connection *redis.Client = nil
)

type CacheManager struct {
	Client *redis.Client
}

func InitBillboardsRedis() {
	config := static.GetCoshubGlobalConfigurations()
	redis_host = config.Redis.Host
	redis_port = config.Redis.Port
	redis_pass = config.Redis.Pass
	redis_connection = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redis_host, redis_port),
		Password: redis_pass,
		DB:       0,
	})
}

func NewCacheManager() *CacheManager {
	if _, err := redis_connection.Ping(context.Background()).Result(); err != nil {
		redis_connection = redis.NewClient(&redis.Options{
			Addr:     redis_host,
			Password: redis_pass,
			DB:       0,
		})
	}
	return &CacheManager{
		Client: redis_connection,
	}
}

func (client *CacheManager) GetRedisKey(namespace string, key string) string {
	return "billboards_web_" + namespace + "_" + key
}

func (client *CacheManager) SetNX(key string, value interface{}, expiration time.Duration) error {
	set_success, err := client.Client.SetNX(context.Background(), key, value, expiration).Result()
	if err != nil {
		return err
	}
	if !set_success {
		return ErrCacheUsedAlready
	}
	return nil
}

func (client *CacheManager) Del(keys ...string) error {
	_, err := client.Client.Del(context.Background(), keys...).Result()
	return err
}

var (
	ErrCacheAdapterNotInitialized = errors.New("cache adapter not initialized")
	ErrCacheKeyNotFount           = errors.New("key not found")
	ErrCacheUsedAlready           = errors.New("cache locked already")
)

func CacheGetGeneric[T any](cache *CacheManager, key interface{}, prefix ...string) (*T, error) {
	if cache == nil || cache.Client == nil {
		return nil, ErrCacheAdapterNotInitialized
	}

	var result T
	// fetch T type name by reflection
	type_name := reflect.TypeOf(result).Name()

	if len(prefix) > 0 {
		type_name = prefix[0] + "_" + type_name
	}

	// check key type, if not string
	key = typ.InterfaceToString(key)
	if key == "" {
		return nil, errors.New("key type not supported")
	}

	real_key := cache.GetRedisKey("generic::"+type_name, key.(string))

	// get from redis
	data, err := cache.Client.Get(
		context.Background(),
		real_key,
	).Result()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		return nil, ErrCacheKeyNotFount
	}

	// unmarshal
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CacheSetGeneric set generic data to redis, different type of data will be stored in different key
// key type must be string, int, float and expire is 0, means never expire
func CacheSetGeneric[T any](cache *CacheManager, key interface{}, value *T, expiration time.Duration, prefix ...string) error {
	if cache == nil || cache.Client == nil {
		return errors.New("cache adapter not initialized")
	}

	// fetch T type name by reflection
	type_name := ""
	// value is a pointer, so we need to get the real type name
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		type_name = reflect.TypeOf(value).Elem().Name()
	} else {
		type_name = reflect.TypeOf(value).Name()
	}

	if len(prefix) > 0 {
		type_name = prefix[0] + "_" + type_name
	}

	// check key type, if not string
	key = typ.InterfaceToString(key)
	if key == "" {
		return errors.New("key type not supported")
	}

	real_key := cache.GetRedisKey("generic::"+type_name, key.(string))

	// marshal
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// set to redis， and expire is 0, means never expire
	err = cache.Client.Set(
		context.Background(),
		real_key,
		data,
		expiration,
	).Err()

	if err != nil {
		return err
	}

	return nil
}

// CacheDelGeneric delete generic data from redis
func CacheDelGeneric[T any](cache *CacheManager, key interface{}, value *T, prefix ...string) error {
	if cache == nil || cache.Client == nil {
		return errors.New("cache adapter not initialized")
	}

	// fetch T type name by reflection
	type_name := ""
	// value is a pointer, so we need to get the real type name
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		type_name = reflect.TypeOf(value).Elem().Name()
	} else {
		type_name = reflect.TypeOf(value).Name()
	}

	if len(prefix) > 0 {
		type_name = prefix[0] + "_" + type_name
	}

	// check key type, if not string
	key = typ.InterfaceToString(key)
	if key == "" {
		return errors.New("key type not supported")
	}

	real_key := cache.GetRedisKey("generic::"+type_name, key.(string))

	// delete from redis
	err := cache.Client.Del(
		context.Background(),
		real_key,
	).Err()

	if err != nil {
		return err
	}

	return nil
}
