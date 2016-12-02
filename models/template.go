package models

import (
  "github.com/jellybean4/gosalt/db"
  "github.com/jellybean4/gosalt/util"
  log "github.com/Sirupsen/logrus"
)

var (
  TemplateCache *Cache
)

func init() {
  TemplateCache = NewTemplateCache()
}

type (
  Template struct {
    Name    string            `json:"name"`
    Configs map[string]string `json:"configs"`
  }

  TemplateHandler struct {
    *DefaultHandler
  }
)

func (h *TemplateHandler) Unmarshal(data []byte) (interface{}, error) {
  v := new(Template)
  if err := util.Unmarshal(data, v); err != nil {
    log.WithFields(log.Fields{
      "data": data,
      "reason": err.Error(),
    }).Error("unmarshal data failed")
    return nil, err
  } else {
    return v, nil
  }
}

func (h *TemplateHandler) Key(v interface{}) (string, bool) {
  if t, ok := v.(*Template); ok {
    log.WithFields(log.Fields{
      "value": v,
      "type": "template",
    }).Error("type assert failed")
    return t.Name, true
  } else {
    return "", false
  }
}

func NewTemplateCache() *Cache {
  dftHandler := &DefaultHandler{table: db.TEMPLATE_TABLE}
  tplHandler := &TemplateHandler{DefaultHandler: dftHandler}
  c := NewCache(string(db.TEMPLATE_TABLE), tplHandler, true, false)
  return c
}

func (t *Template) syncFile() {

}

func (t *Template) delFile() {

}
