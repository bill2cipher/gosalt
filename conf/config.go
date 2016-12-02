package conf

import (
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/gosalt/models"
  "github.com/jellybean4/gosalt/util"
  "github.com/spf13/viper"
)

func Add(svr *models.Server) error {
  if err := models.ServerCache.Set(svr.Name, svr); err != nil {
    return err
  } else if rslt := svr.InitConfig(); rslt.Error != nil {
    return rslt.Error
  } else if rslt := svr.SyncConfig(); rslt.Error != nil {
    return rslt.Error
  }
  return nil
}

func Del(name string) error {
  err := models.ServerCache.Delete(name);
  if err != nil {
    log.WithFields(log.Fields{
      "name": name,
      "reason": err.Error(),
    }).Error("delete svr failed")
    return err
  }
  return nil
}

func Config(version string, servers string) *util.ExecResult {
  confScript := viper.GetString(util.ROOT_DIR) + "/" +
      viper.GetString(util.SYNC_SCRIPT)
  return util.ExecScript(confScript)
}
