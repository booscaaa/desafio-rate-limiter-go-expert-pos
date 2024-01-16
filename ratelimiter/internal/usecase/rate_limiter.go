package usecase

import (
	"context"
	"desafio-rate-limiter-go-expert-pos/ratelimiter/internal/entity"
	"desafio-rate-limiter-go-expert-pos/ratelimiter/internal/infra/inmemory"
	"desafio-rate-limiter-go-expert-pos/ratelimiter/internal/infra/rdb"
	"time"

	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	defaultConfig = entity.Config{}
)

func ConfigLimiter(database entity.DatabaseRepository) {
	ctx := context.Background()
	for _, token := range defaultConfig.Limiter.Tokens {
		rateLimiterInfo := entity.RateLimiterInfo{
			Key:      token.Token,
			Requests: token.Requests,
			Every:    token.Every,
		}

		if err := database.Create(ctx, rateLimiterInfo, 0); err != nil {
			log.Println(err)
		}
	}

	for _, ip := range defaultConfig.Limiter.IPS {
		rateLimiterInfo := entity.RateLimiterInfo{
			Key:      ip.IP,
			Requests: ip.Requests,
			Every:    ip.Every,
		}

		if err := database.Create(ctx, rateLimiterInfo, 0); err != nil {
			log.Println(err)
		}
	}
}

func LoadConfig() {
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}
}

func GetStorage() entity.DatabaseRepository {
	if defaultConfig.Limiter.Database.InMemory {
		return inmemory.NewDatabaseRepository()
	}

	if defaultConfig.Limiter.Database.Redis {
		database := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		return rdb.NewDatabaseRepository(database)
	}

	panic("storage must be provide (inMemory/redis or implement yourself)")
}

func CheckLimit(ctx context.Context, database entity.DatabaseRepository, key string) bool {
	config, err := database.Read(ctx, key)

	fmt.Println(key)

	if err != nil {
		database.Create(ctx, entity.RateLimiterInfo{
			Key:      key,
			Requests: defaultConfig.Limiter.Default.Requests,
			Every:    defaultConfig.Limiter.Default.Every,
		}, 0)

		return CheckLimit(ctx, database, key)
	}

	limiter, err := database.CheckLimit(ctx, key, config.Requests, time.Duration(config.Every)*time.Second)

	if err != nil {
		log.Println(err)
		return false
	}

	return limiter
}
