package controller

import (
  "net/http"
)

import (
  log "github.com/Sirupsen/logrus"
  "github.com/gin-gonic/gin"
  . "github.com/jellybean4/gosalt/mesg"
  "github.com/jellybean4/gosalt/release"
  "github.com/jellybean4/gosalt/util"
)

func Release(c *gin.Context) {
  req, rep := new(ReleaseReq), new(ReleaseRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "release",
      "reason": err.Error(),
    }).Error("read request release data failed")
    rep.Code = 1
    rep.Mesg = err.Error()
  } else if result := release.Release(req.Version, req.Types...); result.Error != nil {
    log.WithFields(log.Fields{
      "action":  "release",
      "request": req,
      "reason":  result.Error.Error(),
      "stderr":  result.Stderr.String(),
    }).Error("execute release failed")
    rep.Code = 2
    rep.Mesg = result.Stderr.String()
  } else {
    log.WithFields(log.Fields{
      "action": "release",
      "request": req,
      "stdout": result.Stdout.String(),
    }).Info("execute release success")
    rep.Code = 0
    rep.Mesg = ""
  }
  c.JSON(http.StatusOK, rep)
}
