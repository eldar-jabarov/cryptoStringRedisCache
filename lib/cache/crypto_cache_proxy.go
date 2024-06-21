package cache

import (
	"cryptoStringRedisCache/internal/infrastructure/service"
	"fmt"
	"log"
)

type CryptoCacheProxy struct {
	obj   service.Crypter
	cache Cacher
}

func NewServiceCrypto(obj service.Crypter, cache Cacher) service.Crypter {
	return &CryptoCacheProxy{obj: obj, cache: cache}
}

func (c *CryptoCacheProxy) MD5(s string) (string, error) {
	key := fmt.Sprintf("MD5:%s", s)
	res, err := c.cache.Get(key)
	if err != nil {
		log.Printf("cache.CryptoCacheProxy.MD5(%s).cache.Get err: %s", s, err)

		res, err = c.obj.MD5(s)
		if err != nil {
			return "", err
		}

		err = c.cache.Set(key, res)
		if err != nil {
			log.Printf("cache.CryptoCacheProxy.MD5(%s).cache.Set err: %s", s, err)
		}

		return res, nil
	}

	return res, err
}

func (c *CryptoCacheProxy) SHA256(s string) (string, error) {
	key := fmt.Sprintf("SHA256:%s", s)
	res, err := c.cache.Get(key)
	if err != nil {
		log.Printf("cache.CryptoCacheProxy.SHA256(%s).cache.Get err: %s", s, err)

		res, err = c.obj.SHA256(s)
		if err != nil {
			return "", err
		}

		err = c.cache.Set(key, res)
		if err != nil {
			log.Printf("cache.CryptoCacheProxy.SHA256(%s).cache.Set err: %s", s, err)
		}

		return res, nil
	}

	return res, err
}
