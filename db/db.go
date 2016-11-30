package db

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
)

// db is used to store dev info
var gosaltdb *bolt.DB

func init() {
	options := &bolt.Options{Timeout: 1 * time.Second}
  dbFile := viper.GetString(util.DB_FILE)
	if db, err := bolt.Open(dbFile, 0600, options); err != nil {
		log.WithFields(log.Fields{
			"dbfile": dbFile,
			"reason": err.Error(),
		}).Fatal("open dbfile failed")
	} else {
		log.WithFields(log.Fields{
			"dbfile": dbFile,
		}).Info("open dbfile success")
		gosaltdb = db
	}
	createTables()
}

func createTables() {
	for tb := range TABLES {
		err := gosaltdb.Update(func(tx *bolt.Tx) error {
			if _, err := tx.CreateBucketIfNotExists([]byte(tb)); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.WithFields(log.Fields{
				"table":  tb,
				"reason": err.Error(),
			}).Fatal("create table failed")
		}
	}
}

func Get(table []byte, key []byte) ([]byte, error) {
	var data []byte
	err := gosaltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		data = b.Get(key)
		return nil
	})
	if err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"key":    key,
			"reason": err.Error(),
		}).Error("db fetch data failed")
	}
	return data, err
}

func Set(table []byte, key []byte, value []byte) error {
	err := gosaltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		return b.Put(key, value)
	})
	if err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"key":    key,
			"value":  value,
			"reason": err.Error(),
		}).Error("store value to db failed")
	}
	return err
}

func Unset(table []byte, key []byte) error {
	err := gosaltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		return b.Delete(key)
	})
	if err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"key":    key,
			"reason": err.Error(),
		}).Error("delete db data failed")
	}
	return err
}

func All(table []byte) ([][]byte, error) {
	var datas [][]byte
	err := gosaltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		return b.ForEach(func(k, v []byte) error {
			datas = append(datas, v)
			return nil
		})
	})
	if err != nil {
		log.WithFields(log.Fields{
			"table":  table,
			"reason": err.Error(),
		}).Error("fetch all db data failed")
	}
	return datas, err
}
