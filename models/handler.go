package models

import (
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/gosalt/db"
  "github.com/jellybean4/gosalt/util"
)

type DefaultHandler struct {
  table []byte
}

func (h *DefaultHandler) Marshal(value interface{}) ([]byte, error) {
  return util.Marshal(value)
}

func (h *DefaultHandler) Get(key string) ([]byte, bool) {
  if v, err := db.Get(h.table, []byte(key)); err != nil {
    log.WithFields(log.Fields{
      "name": h.table,
      "key": key,
      "reason": err.Error(),
    }).Error("handler fetch boltdb data failed")
    return nil, false
  } else {
    log.WithFields(log.Fields{
      "name": h.table,
      "key": key,
      "value": v,
    }).Info("handler fetch boltdb data success")
    return v, true
  }
}

func (h *DefaultHandler) Delete(key string) error {
  if err := db.Unset(h.table, []byte(key)); err != nil {
    log.WithFields(log.Fields{
      "name": h.table,
      "key":key,
      "reason": err.Error(),
    }).Error("handler delete boltdb data failed")
    return err
  } else {
    log.WithFields(log.Fields{
      "name": h.table,
      "key": key,
    }).Info("handler delete boltdb data success")
    return nil
  }
}

func (h *DefaultHandler) Set(key string, value []byte) error {
  if err := db.Set(h.table, []byte(key), value); err != nil {
    log.WithFields(log.Fields{
      "name": h.table,
      "key":key,
      "value": value,
      "reason": err.Error(),
    }).Error("handler set boltdb data failed")
    return err
  } else {
    log.WithFields(log.Fields{
      "name": h.table,
      "key": key,
      "value": value,
    }).Info("handler set boltdb data success")
    return nil
  }
}

func (h *DefaultHandler) All() [][]byte {
  if data, err := db.All(h.table); err != nil {
    log.WithFields(log.Fields{
      "name": h.table,
      "reason": err.Error(),
    }).Error("load all from boltdb failed")
    return nil
  } else {
    log.WithFields(log.Fields{
      "name": h.table,
    }).Info("load all from boltdb success")
    return data
  }
}
