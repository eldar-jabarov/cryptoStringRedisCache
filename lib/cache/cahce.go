package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

var NotFoundErr = errors.New("cache did not find data")

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) Cacher {
	return &Cache{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	return c.client.Set(context.Background(), key, value, 0).Err()
}

func (c *Cache) Get(key string) (string, error) {
	s := c.client.Get(context.Background(), key)
	res, err := s.Result()
	if errors.Is(err, redis.Nil) {
		res, err = "", NotFoundErr
	}
	return res, err
}
