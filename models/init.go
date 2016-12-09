package models

import (
  "github.com/jellybean4/order_init"
  "github.com/jellybean4/gosalt/util"
)

func init() {
  order.RegisteFunc(util.MODEL_SVR_INIT, initServer, util.DB_INIT, util.LOG_INIT)
  order.RegisteFunc(util.MODEL_TEMPL_INIT, initTempl, util.DB_INIT, util.LOG_INIT)
  order.RegisteFunc(util.MODEL_TEMPL_LIB_INIT, initTemplLib, util.LOG_INIT)
  order.RegisteFunc(util.MODEL_SVR_LIB_INIT, initServerLib, util.LOG_INIT)
}
