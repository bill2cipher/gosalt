package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
)

var (
	TemplateCache *Cache
)

func initTempl() {
  if c, err := NewTemplateCache(); err != nil {
    log.WithFields(log.Fields{
      "table": db.SERVER_TABLE,
      "reason": err.Error(),
    }).Fatal(util.CACHE_INIT_DATA_LOG)
  } else {
    TemplateCache = c
  }
}

type (
	Template struct {
		Name    string            `json:"name"`
		Env     string            `json:"env"`
		Version string            `json:"version"`
		Config  map[string]string `json:"config"`
	}

	TemplateHandler struct {
		*DefaultHandler
	}
)

func (h *TemplateHandler) Unmarshal(data []byte) (interface{}, error) {
	v := new(Template)
	if err := util.Unmarshal(data, v); err != nil {
		log.WithFields(log.Fields{
			"data":   data,
			"reason": err.Error(),
		}).Error(util.JSON_UNMARSHAL_LOG)
		return nil, err
	} else {
		return v, nil
	}
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
	dftHandler := &DefaultHandler{table: db.TEMPLATE_TABLE}
	tplHandler := &TemplateHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.TEMPLATE_TABLE), tplHandler, true, false)
}
