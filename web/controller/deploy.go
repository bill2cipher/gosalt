package controller

import (
  "net/http"
)

import (
  "github.com/gin-gonic/gin"
  log "github.com/Sirupsen/logrus"
  . "github.com/jellybean4/gosalt/mesg"
  "github.com/jellybean4/gosalt/util"
  "github.com/jellybean4/gosalt/deploy"
)

func Deploy(c *gin.Context) {
  req, rep := new(DeployReq), new(DeployRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "deploy",
      "reason": err.Error(),
    }).Error("read request deploy data failed")
    rep.Code = 1
    rep.Mesg = err.Error()
  } else if result := deploy.Deploy(req.Version, req.Server); result.Error != nil {
    log.WithFields(log.Fields{
      "action": "deploy",
      "request": req,
      "reason": result.Error.Error(),
      "stderr": result.Stderr.String(),
    }).Error("execute deploy failed")
    rep.Code = 2
    rep.Mesg = err.Error()
  } else {
    log.WithFields(log.Fields{
      "action": "release",
      "request": req,
      "stdout": result.Stdout.String(),
    }).Info("execute deploy success")
    rep.Code = 0
    rep.Mesg = ""
  }
  c.JSON(http.StatusOK, rep)
}
