package Redis

import (
	"chat/Config"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisService 負責 Redis 操作
type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	config, err := Config.LoadConfig()
	if err != nil {
		panic("config loading failed")
	}
	return &RedisService{client: NewRedisClient(&config.Redis)}
}

// NewRedisClient 建立 Redis 客戶端，使用來自 YAML 的配置
func NewRedisClient(config *Config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         config.Address,
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
}

// GetValue 獲取 Redis 中的數值
func (r *RedisService) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// SetValue 設定 Redis 中的 key-value 值
func (r *RedisService) SetValue(ctx context.Context, key string, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err() // 0 表示不設置過期時間
	if err != nil {
		return err
	}
	return nil
}

// ListPush 向 List 中推入元素
func (r *RedisService) ListPush(ctx context.Context, key string, values ...string) error {
	err := r.client.LPush(ctx, key, values).Err()
	if err != nil {
		return err
	}
	return nil
}

// ListRange 獲取 List 中的元素
func (r *RedisService) ListRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	result, err := r.client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HashSet 設置 Hash 中的字段值
func (r *RedisService) HashSet(ctx context.Context, key, field, value string) error {
	err := r.client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// HashGet 獲取 Hash 中的字段值
func (r *RedisService) HashGet(ctx context.Context, key, field string) (string, error) {
	result, err := r.client.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("field %s does not exist in key %s", field, key)
	} else if err != nil {
		return "", err
	}
	return result, nil
}
