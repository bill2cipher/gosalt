package db

import (
  "github.com/jellybean4/order_init"
  "github.com/jellybean4/gosalt/util"
  log "github.com/Sirupsen/logrus"
  "gopkg.in/mgo.v2/bson"
  "fmt"
  "gopkg.in/mgo.v2"
)


//TODO: db pool
func init() {
  order.RegisteFunc(util.DB_INIT, initDB, util.CONFIG_INIT, util.LOG_INIT)
}

func initSession() {
  connStr := fmt.Sprintf("mongodb://%s:%s@%s/%s",
    util.GetConfig(util.DB_USER), util.GetConfig(util.DB_PASS),
    util.GetConfig(util.DB_HOST), util.GetConfig(util.DB_NAME))

  sess, err := mgo.Dial(connStr)
  if err != nil {
    log.WithFields(log.Fields{
      "conn":   connStr,
      "reason": err.Error(),
    }).Fatal(util.DB_CONN_LOG)
  } else {
    sess.SetMode(mgo.Monotonic, true)
    GLOBAL_SESSION = sess
  }
}

func initDB() {
  initSession()

  coll, sess, err := getCollection(COUNTER_TABLE)
  if err != nil {
    log.WithFields(log.Fields{
      "table": COUNTER_TABLE,
      "reason": err.Error(),
    }).Fatal(util.DB_CONN_LOG)
  }
  defer sess.Close()

  for _, table := range TABLES {
    query := coll.Find(bson.M{"_id": table})
    if cnt, err := query.Count(); err != nil {
      log.WithFields(log.Fields{
        "table": COUNTER_TABLE,
        "reason": err.Error(),
      }).Fatal(util.DB_FETCH_LOG)
    } else if cnt != 0 {
      cnter := new(Counter)
      if err := query.One(cnter); err != nil {
        log.WithFields(log.Fields{
          "table": COUNTER_TABLE,
          "reason": err.Error(),
        }).Fatal(util.DB_FETCH_LOG)
      } else {
        startCounter(table, cnter.Seq)
      }
    } else if coll.Insert(bson.M{"_id": table, "seq": 1}); err != nil {
      log.WithFields(log.Fields{
        "table": COUNTER_TABLE,
        "reason": err.Error(),
      }).Fatal(util.DB_STORE_LOG)
    } else {
      startCounter(table, 1)
    }
  }
}

