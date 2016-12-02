package models

import (
	"time"
)

import (
	log "github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
)

const (
	EXPIRE_INTERVAL  = 5 * time.Minute
	CLEANUP_INTERVAL = 30 * time.Second
)

type CacheHandler interface {
	Marshal(value interface{}) ([]byte, error)  // marshal go obj
	Unmarshal(data []byte) (interface{}, error) // unmarshal go obj

	Key(v interface{}) (string, bool)   // parse key from obj
	Get(key string) ([]byte, bool)      // Get data
	Set(key string, value []byte) error // set data
	All() [][]byte                      // load all data
	Delete(key string) error            // delete data
}

type Cache struct {
	name string
	c    *cache.Cache
	h    CacheHandler
}

func NewCache(name string, h CacheHandler, load, expire bool) *Cache {
	c := new(Cache)
	c.name = name
	c.h = h

	if !load && !expire {
		c.c = cache.New(cache.NoExpiration, CLEANUP_INTERVAL)
	} else if !load && expire {
		c.c = cache.New(EXPIRE_INTERVAL, CLEANUP_INTERVAL)
	} else if load && expire {
		items := c.loadAll(expire)
		c.c = cache.NewFrom(EXPIRE_INTERVAL, CLEANUP_INTERVAL, items)
	} else if load && !expire {
		items := c.loadAll(expire)
		c.c = cache.NewFrom(cache.NoExpiration, CLEANUP_INTERVAL, items)
	}
	return c
}

func (c *Cache) loadAll(expire bool) (items map[string]cache.Item) {
	all, now := c.h.All(), time.Now().UnixNano()
	for _, d := range all {
		if v, err := c.h.Unmarshal(d); err != nil {
			continue
		} else if k, ok := c.h.Key(v); !ok {
			continue
		} else if !expire {
			items[k] = cache.Item{Object: v, Expiration: 0}
		} else {
			items[k] = cache.Item{Object: v, Expiration: now + int64(EXPIRE_INTERVAL)}
		}
	}
	return items
}

func (c *Cache) Get(key string) interface{} {
	if v, ok := c.c.Get(key); ok {
		log.WithFields(log.Fields{
			"name":  c.name,
			"key":   key,
			"value": v,
		}).Info("cache get data success")
		return v
	} else if d, ok := c.h.Get(key); !ok {
		log.WithFields(log.Fields{
			"name":   c.name,
			"key":    key,
			"reason": "data not exist",
		}).Warn("cache fetch persiste data failed")
		return nil
	} else if uv, err := c.h.Unmarshal(d); err != nil {
		log.WithFields(log.Fields{
			"name":   c.name,
			"key":    key,
			"reason": err.Error(),
		}).Error("cache unmarshal persiste data failed")
		return nil
	} else {
		log.WithFields(log.Fields{
			"name": c.name,
			"key":  key,
		}).Info("cache fetch persiste data success")
		c.c.Set(key, uv, cache.NoExpiration)
		return uv
	}
}

func (c *Cache) Delete(key string) error {
	c.c.Delete(key)
	if err := c.h.Delete(key); err != nil {
		log.WithFields(log.Fields{
			"name":   c.name,
			"key":    key,
			"reason": err.Error(),
		}).Error("cache delete data failed")
		return err
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) error {
	if data, err := c.h.Marshal(value); err != nil {
		return err
	} else if err := c.h.Set(key, data); err != nil {
		return err
	} else {
		c.c.Set(key, value, cache.NoExpiration)
		return nil
	}
}
