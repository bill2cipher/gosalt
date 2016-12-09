package util

import (
  "github.com/jellybean4/order_init"
  log "github.com/Sirupsen/logrus"
)


func init() {
  order.RegisteFunc(UTIL_INIT, initUtil)
  order.RegisteFunc(LOG_INIT, initLog)
}

func initLog() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetLevel(log.DebugLevel)
}
