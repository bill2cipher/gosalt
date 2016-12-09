package models

import (
	"github.com/jellybean4/gosalt/db"
	"reflect"
)

type DefaultHandler struct {
	table string
	t     reflect.Type
}

func (h *DefaultHandler) Get(key string) (interface{}, error) {
	if v, err := db.Get(h.table, key, h.t); err != nil {
		return nil, err
	} else {
		return v, nil
	}
}

func (h *DefaultHandler) Delete(key string) error {
	if err := db.Unset(h.table, key); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *DefaultHandler) Set(key string, value interface{}) error {
	if err := db.Set(h.table, key, value); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *DefaultHandler) All() ([]interface{}, error) {
	if datas, err := db.All(h.table, h.t); err != nil {
		return nil, err
	} else {
		return datas, nil
	}
}
