package controller

import (
  "net/http"
)

import (
  "github.com/gin-gonic/gin"
  log "github.com/Sirupsen/logrus"
  . "github.com/jellybean4/gosalt/mesg"
  "github.com/jellybean4/gosalt/util"
  "github.com/jellybean4/gosalt/module"
)

func Deploy(c *gin.Context) {
  req, rep := new(DeployReq), new(DeployRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "deploy",
      "reason": err.Error(),
    }).Error(util.GIN_PARSE_REQUEST_LOG)
    rep.Code = util.GIN_PARSE_REQUEST_CODE
    rep.Mesg = err.Error()
  } else if result := module.Deploy(req.Version, req.Server); result.Error != nil {
    log.WithFields(log.Fields{
      "action": "deploy",
      "request": req,
      "reason": result.Error.Error(),
      "stderr": result.Stderr.String(),
    }).Error(util.SCRIPT_EXECUTE_LOG)
    rep.Code = util.SCRIPT_EXECUTE_CODE
    rep.Mesg = err.Error()
  } else {
    log.WithFields(log.Fields{
      "action": "release",
      "request": req,
      "stdout": result.Stdout.String(),
    }).Info("execute deploy success")
    rep.Code = util.SUCCESS_CODE
    rep.Mesg = util.SUCCESS_MESG
  }
  c.JSON(http.StatusOK, rep)
}
