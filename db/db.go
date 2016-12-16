package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/util"
)

import (
	"gopkg.in/mgo.v2"
	"reflect"
)

const (
	EMPTY_ID = "empty"
)

var (
  GLOBAL_SESSION *mgo.Session
)

func getCollection(table string) (*mgo.Collection, *mgo.Session, error) {
  sess := GLOBAL_SESSION.Clone()
	return sess.DB(util.GetConfig(util.DB_NAME)).C(table), sess, nil
}

func Get(table string, key string, t reflect.Type) (interface{}, error) {
	coll, sess, err := getCollection(table)
	if err != nil {
		return nil, err
	}
  defer sess.Close()

	value := reflect.New(t).Interface()
	if err := coll.FindId(key).One(value); err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"key":    key,
			"reason": err.Error(),
		}).Error(util.DB_FETCH_LOG)
		return nil, err
	} else {
		return value, nil
	}
}

func Set(table string, key string, value interface{}) error {
	coll, sess, err := getCollection(table)
	if err != nil {
		return err
	}
  defer sess.Close()

	if _, err := coll.UpsertId(key, value); err != nil {
    log.WithFields(log.Fields{
      "table":  table,
      "key":    key,
      "value":  value,
      "reason": err.Error(),
    }).Error(util.DB_STORE_LOG)
    return err
  } else {
    log.WithFields(log.Fields{
      "table": table,
      "key":   key,
      "value": value,
    }).Info("db store success")
    return nil
  }
}
  
func Unset(table, key string) error {
	coll, sess, err := getCollection(table)
	if err != nil {
		return err
	}
  defer sess.Close()

	if err := coll.RemoveId(key); err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"key":    key,
			"reason": err.Error(),
		}).Error(util.DB_DELETE_LOG)
		return err
	} else {
		log.WithFields(log.Fields{
			"table": table,
			"key":   key,
		}).Info("db delete success")
		return nil
	}
}

func All(table string, t reflect.Type) ([]interface{}, error) {
	coll, sess, err := getCollection(table)
	if err != nil {
		return nil, err
	}
  defer sess.Close()

	query := coll.Find(nil)
	iter := query.Iter()
	cnt, err := query.Count()
	if err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"reason": err.Error(),
		}).Error(util.DB_FETCH_LOG)
		return nil, err
	}
	result, i := make([]interface{}, cnt), 0

	for i < cnt {
		value := reflect.New(t).Interface()
		if !iter.Next(value) {
			break
		}
		result[i] = value
		i++
	}

	if i != cnt {
		result = result[0:i]
	}

	if err := iter.Close(); err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"reason": err.Error(),
		}).Error(util.DB_FETCH_LOG)
		return nil, err
	} else {
		return result, nil
	}
}

