package models

import (
	"time"
)

import (
	"github.com/patrickmn/go-cache"
)

const (
	EXPIRE_INTERVAL  = 5 * time.Minute
	CLEANUP_INTERVAL = 1 * time.Minute
)

type CacheHandler interface {
	Key(v interface{}) (string, bool)        // parse key from obj
	Get(key string) (interface{}, error)      // Get data
	Set(key string, value interface{}) error // set data
	All() ([]interface{}, error)             // load all data
	Delete(key string) error                 // delete data
}

type Cache struct {
	name   string
	c      *cache.Cache
	h      CacheHandler
}

func NewCache(name string, h CacheHandler, load, expire bool) (*Cache, error) {
	c := new(Cache)
	c.name = name
	c.h = h

	var (
		items map[string]cache.Item
		err   error
	)
	if load {
		items, err = c.loadAll(expire)
		if err != nil {
			return nil, err
		}
	}

	if !load && !expire {
		c.c = cache.New(cache.NoExpiration, CLEANUP_INTERVAL)
	} else if !load && expire {
		c.c = cache.New(EXPIRE_INTERVAL, CLEANUP_INTERVAL)
	} else if load && expire {
		c.c = cache.NewFrom(EXPIRE_INTERVAL, CLEANUP_INTERVAL, items)
	} else if load && !expire {
		c.c = cache.NewFrom(cache.NoExpiration, CLEANUP_INTERVAL, items)
	}
	return c, nil
}

func (c *Cache) loadAll(expire bool) (items map[string]cache.Item, err error) {
	all, err := c.h.All()
	now := time.Now().UnixNano()
	if err != nil {
		return nil, err
	}
	for _, v := range all {
		if k, ok := c.h.Key(v); !ok {
			continue
		} else if !expire {
			items[k] = cache.Item{Object: v, Expiration: 0}
		} else {
			items[k] = cache.Item{Object: v, Expiration: now + int64(EXPIRE_INTERVAL)}
		}
	}
	return items, nil
}

func (c *Cache) Get(key string) interface{} {
	if v, ok := c.c.Get(key); ok {
		return v
	} else if v, err := c.h.Get(key); err != nil {
		return nil
	} else {
		c.c.Set(key, v, cache.DefaultExpiration)
		return v
	}
}

func (c *Cache) Delete(key string) error {
	c.c.Delete(key)
	if err := c.h.Delete(key); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) error {
	if err := c.h.Set(key, value); err != nil {
		return err
	} else {
		c.c.Set(key, value, cache.NoExpiration)
		return nil
	}
}
