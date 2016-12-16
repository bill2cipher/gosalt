package test

import (
  "testing"
  "github.com/jellybean4/gosalt/models"
  "github.com/jellybean4/gosalt/db"
  "reflect"
)

func TestServer(t *testing.T) {
  svr := &models.Server{"test", "test", 22, "good", "passwd", "dev", "trunk"}
  if err := models.SetServer(svr); err != nil {
    t.Errorf("set server %v failed, reason %s", svr, err.Error())
  }

  if val := models.ServerCache.Get("test"); val == nil {
    t.Errorf("get cache server failed, not exist")
  } else if s, ok := models.ServerCache.Get("test").(*models.Server); !ok {
    t.Errorf("get cache server failed, type assert failed %v", reflect.TypeOf(val))
  } else if svr.Name != s.Name {
    t.Errorf("get svr %v failed", val)
  }

  if v, err := db.Get(db.SERVER_TABLE, "test", reflect.TypeOf(models.Server{})); err != nil {
    t.Errorf("get db server failed, reason %s", err.Error())
  } else if s, ok := v.(*models.Server); !ok {
    t.Errorf("get db server type failed %v", v)
  } else if s.Name != svr.Name {
    t.Errorf("get db server content failed")
  }
}
