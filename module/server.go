package module

import (
  "errors"
)

import (
  "github.com/jellybean4/gosalt/models"
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/gosalt/util"
)

func AddServer(svr *models.Server) error {
  if v := models.ServerCache.Get(svr.Name); v == nil {
    return addNewServer(svr)
  } else if _, ok := v.(*models.Server); !ok {
    log.WithFields(log.Fields{
      "type": "*models.Server",
      "value": v,
      "reason": "type assert failed",
    }).Error(util.TYPE_ASSERT_LOG)
    return errors.New(util.TYPE_ASSERT_LOG)
  } else {
    return editServer(svr)
  }
}

func addNewServer(svr *models.Server) error {
  return nil
}

func editServer(svr *models.Server) error {
  if err := models.SetServer(svr); err != nil {
    return err
  } else {
    return nil
  }
}