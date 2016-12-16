package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
	"reflect"
  "errors"
)

var (
	ConfigsCache *Cache
)

type (
	Config struct {
		ID      string            `bson:"_id"`
		Name    string            `bson:"name"`
		Version string            `bson:"version"`
		Env     string            `bson:"env"`
		Config  map[string]string `bson:"config"`
	}

	ConfigHandler struct {
		*DefaultHandler
	}
)

func (h *ConfigHandler) SetKey(v interface{}, key string) error {
  if c, ok := v.(*Config); !ok {
    log.WithFields(log.Fields{
      "value": v,
      "type": "models.Config",
    }).Error(util.TYPE_ASSERT_LOG)
    return errors.New(util.TYPE_ASSERT_LOG)
  } else {
    c.ID = key
    return nil
  }
}

func (h *ConfigHandler) Key(v interface{}) (string, bool) {
	if t, ok := v.(*Template); ok {
		return t.Name, true
	} else {
		log.WithFields(log.Fields{
			"value": v,
			"type":  "configs",
		}).Error(util.TYPE_ASSERT_LOG)
		return "", false
	}
}

func NewConfigCache() (*Cache, error) {
	dftHandler := &DefaultHandler{table: db.CONFIGS_TABLE, t: reflect.TypeOf(Config{})}
	cfgHandler := &ConfigHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.CONFIGS_TABLE), cfgHandler, false, true)
}
