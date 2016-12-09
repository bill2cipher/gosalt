package controller

import (
  "net/http"
)

import (
  log "github.com/Sirupsen/logrus"
  "github.com/gin-gonic/gin"
  . "github.com/jellybean4/gosalt/mesg"
  "github.com/jellybean4/gosalt/module"
  "github.com/jellybean4/gosalt/util"
)

func Release(c *gin.Context) {
  req, rep := new(ReleaseReq), new(ReleaseRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "release",
      "reason": err.Error(),
    }).Error(util.GIN_PARSE_REQUEST_LOG)
    rep.Code = util.GIN_PARSE_REQUEST_CODE
    rep.Mesg = err.Error()
  } else if result := module.Release(req.Version, req.Types...); result.Error != nil {
    log.WithFields(log.Fields{
      "action":  "release",
      "request": req,
      "reason":  result.Error.Error(),
      "stderr":  result.Stderr.String(),
    }).Error(util.SCRIPT_EXECUTE_LOG)
    rep.Code = util.SCRIPT_EXECUTE_CODE
    rep.Mesg = result.Stderr.String()
  } else {
    log.WithFields(log.Fields{
      "action": "release",
      "request": req,
      "stdout": result.Stdout.String(),
    }).Info("execute release success")
    rep.Code = util.SUCCESS_CODE
    rep.Mesg = util.SUCCESS_MESG
  }
  c.JSON(http.StatusOK, rep)
}
