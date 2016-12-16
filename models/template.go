package models

import (
  "errors"
)

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
  "reflect"
)

var (
	TemplateCache *Cache
)

func initTempl() {
	if c, err := NewTemplateCache(); err != nil {
		log.WithFields(log.Fields{
			"table":  db.SERVER_TABLE,
			"reason": err.Error(),
		}).Fatal(util.CACHE_INIT_DATA_LOG)
	} else {
		TemplateCache = c
	}
}

type (
	Template struct {
		Name    string            `json:"name"    bson:"_id"`
		Env     string            `json:"env"     bson:"env"`
		Version string            `json:"version" bson:"version"`
		Config  map[string]string `json:"config"  bson:"config"`
	}

	TemplateHandler struct {
		*DefaultHandler
	}
)

func (h *TemplateHandler) SetKey(v interface{}, key string) error {
  return errors.New("server model not support")
}

func (h *TemplateHandler) Key(v interface{}) (string, bool) {
	if t, ok := v.(*Template); ok {
		return t.Name, true
	} else {
		log.WithFields(log.Fields{
			"value": v,
			"type":  "template",
		}).Error(util.TYPE_ASSERT_LOG)
		return "", false
	}
}

func NewTemplateCache() (*Cache, error) {
	dftHandler := &DefaultHandler{table: db.TEMPLATE_TABLE, t : reflect.TypeOf(Template{})}
	tplHandler := &TemplateHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.TEMPLATE_TABLE), tplHandler, true, false)
}
