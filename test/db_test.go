package test

import (
	"testing"
  "github.com/jellybean4/gosalt/db"
)

func TestDB(t *testing.T) {
	type TS struct {
		ID   string            `bson:"_id"`
		Name string            `bson:"name"`
		Addr string            `bson:"addr"`
		Prop map[string]string `bson:"prop"`
	}

  startInit(t)
  ts := TS{Name: "bad", Addr: "Think",
		Prop: map[string]string{"o": "p", "q": "s", "b": "d"}}
	db.Set("test", db.EMPTY_ID, &ts)
}
