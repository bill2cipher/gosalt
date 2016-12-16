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

  ts := TS{ID : db.EMPTY_ID, Name: "bad", Addr: "Think",
		Prop: map[string]string{"o": "p", "q": "s", "b": "d"}}
	if err := db.Set(db.SERVER_TABLE, db.EMPTY_ID, &ts); err != nil {
    t.Error(err.Error())
  }
}


func TestCounter(t *testing.T) {
  for _, table := range db.TABLES {
    for i := 0; i < 10; i++ {
      if _, err := db.GetNextID(table); err != nil {
        t.Error(err.Error())
      }
    }
  }
}
