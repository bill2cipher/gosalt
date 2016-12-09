package db

import (
  "github.com/jellybean4/order_init"
  "github.com/jellybean4/gosalt/util"
)

func init() {
  order.RegisteFunc(util.DB_INIT, initDB, util.CONFIG_INIT, util.LOG_INIT)
}

func initDB() {

}
