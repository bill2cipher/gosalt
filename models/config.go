package models

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/jellybean4/gosalt/mesg"
	"github.com/jellybean4/gosalt/util"
  "github.com/jellybean4/gosalt/db"
)

var (
	ConfigsCache *Cache
)

type (
	Config map[string]string

	ConfigHandler struct {
		*DefaultHandler
	}
)

func (h *ConfigHandler) Unmarshal(data []byte) (interface{}, error) {
	v := new(Config)
	if err := util.Unmarshal(data, v); err != nil {
		log.WithFields(log.Fields{
			"data":   data,
			"reason": err.Error(),
		}).Error(JSON_UNMARSHAL_LOG)
		return nil, err
	} else {
		return v, nil
	}
}

func (h *ConfigHandler) Key(v interface{}) (string, bool) {
	if t, ok := v.(*Template); ok {
		return t.Name, nil
	} else {
		log.WithFields(log.Fields{
			"value": v,
			"type":  "configs",
		}).Error(TYPE_ASSERT_LOG)
		return "", false
	}
}

func NewConfigCache() *Cache {
  dftHandler := &DefaultHandler{table: db.CONFIGS_TABLE}
  cfgHandler := &ConfigHandler{DefaultHandler: dftHandler}
  c := NewCache(string(db.CONFIGS_TABLE), cfgHandler, false, true)
  return c
}
