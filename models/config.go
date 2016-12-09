package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
  "reflect"
)

var (
	ConfigsCache *Cache
)

type (
	Config struct {
		ID      uint   `gorm:"primary_key;AUTO_INCREMENT;column:name"`
		Name    string `gorm:"column:name"`
		Version string `gorm:"column:version"`
		Config  map[string]string
	}

	ConfigHandler struct {
		*DefaultHandler
	}
)

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
	dftHandler := &DefaultHandler{table: db.CONFIGS_TABLE, t : reflect.TypeOf(Config{})}
	cfgHandler := &ConfigHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.CONFIGS_TABLE), cfgHandler, false, true)
}
