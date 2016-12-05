package models

import (
	log "github.com/Sirupsen/logrus"
  . "github.com/jellybean4/gosalt/mesg"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
)

var (
	TemplateCache *Cache
)

func init() {
	TemplateCache = NewTemplateCache()
}

type (
	Template struct {
		Name   string            `json:"name"`
		Config map[string]string `json:"config"`
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
		}).Error(JSON_UNMARSHAL_LOG)
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
    }).Error(TYPE_ASSERT_LOG)
		return "", false
	}
}

func NewTemplateCache() *Cache {
	dftHandler := &DefaultHandler{table: db.TEMPLATE_TABLE}
	tplHandler := &TemplateHandler{DefaultHandler: dftHandler}
	c := NewCache(string(db.TEMPLATE_TABLE), tplHandler, true, false)
	return c
}

