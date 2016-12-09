package models

import (
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
)


// InitConfig is used to init the server's salt minion config
func InitConfig(svr *Server, args ...string) *util.ExecResult {
  initScript := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.INIT_SCRIPT)
  return util.ExecScript(initScript, args...)
}

// SyncConfig is used to sync config info into minion
func SyncConfig(svr *Server, args ...string) *util.ExecResult {
  syncScript := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.SYNC_SCRIPT)
  return util.ExecScript(syncScript, args...)
}
