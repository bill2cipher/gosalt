package saltstack

import (
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/order_init"
  "github.com/jellybean4/gosalt/util"
)

func init() {
  order.RegisteFunc(util.SALT_INIT, initSalt, util.CONFIG_INIT, util.LOG_INIT)
}

func initSalt() {
  if err := saltLogin(); err != nil {
    log.Fatalf("init salt token failed, reason %s", err.Error())
  } else {
    go startEventReceiver()
    log.Debugf("init salt token %v success", globalToken)
  }
}
