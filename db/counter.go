package db

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/util"
	"gopkg.in/mgo.v2/bson"
)

type Counter struct {
	ID  string `bson:"_id"`
	Seq int    `bson:"seq"`
}

type generator struct {
	input chan bool
	outer chan int
}

var (
	genMap = make(map[string]*generator)
)

func GetNextID(table string) (string, error) {
	if gen, ok := genMap[table]; !ok {
		log.WithFields(log.Fields{
			"table":  table,
			"reason": "genMap item not exist",
		}).Error(util.DB_GEN_ID)
		return EMPTY_ID, errors.New("genMap item not exist")
	} else {
		gen.input <- true
		if id := <-gen.outer; id == 0 {
			return "", errors.New("gen id failed")
		} else {
			return fmt.Sprintf("%d", id), nil
		}
	}
}

func startCounter(table string, next int) {
	input := make(chan bool, 128)
	outer := make(chan int, 256)
	genMap[table] = &generator{input, outer}
	go counterProc(input, outer, table, next)
}

func counterProc(input chan bool, outer chan int, table string, next int) {
	defer func() {
		if err := recover(); err != nil {
			go counterProc(input, outer, table, next)
		}
	}()
	for {
		select {
		case <-input:
			coll, sess, err := getCollection(COUNTER_TABLE)
			if err != nil {
				outer <- 0
				continue
			}
			defer sess.Clone()

			if err := coll.Update(bson.M{"_id": table}, bson.M{"seq": next + 1}); err != nil {
				outer <- 0
			} else {
				next = next + 1
				outer <- next
			}
		}
	}
}
