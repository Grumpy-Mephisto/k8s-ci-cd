package config

import (
	"os-container-project/pkg/utils"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     utils.GetEnv("REDIS_HOST", "localhost").(string),
		Port:     utils.GetEnv("REDIS_PORT", "6379").(string),
		Password: utils.GetEnv("REDIS_PASSWORD", "redis").(string),
		DB:       0,
	}
}

func (r *RedisConfig) NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Password,
		DB:       r.DB,
	})
}
