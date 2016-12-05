package controller

import (
  "github.com/gin-gonic/gin"
  . "github.com/jellybean4/gosalt/mesg"
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/gosalt/util"
  "github.com/jellybean4/gosalt/models"
  "net/http"
)

func SetTemplate(c *gin.Context) {
  req, rep := new(SetTemplateReq), new(SetTemplateRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "template.set",
      "reason": err.Error(),
    }).Error(GIN_PARSE_REQUEST_LOG)
    rep.Code = GIN_PARSE_REQUEST_CODE
    rep.Mesg = err.Error()
  } else {
    templ := new(models.Template)
    templ.Name, templ.Config = req.Name, req.Config
    if err := models.TemplateCache.Set(req.Name, templ); err != nil {
      log.WithFields(log.Fields{
        "action": "template.set",
        "req": req,
        "reason": err.Error(),
      }).Error(CACHE_SET_DATA_LOG)
      rep.Code = CACHE_SET_DATA_CODE
      rep.Mesg = err.Error()
    } else {
      rep.Code, rep.Mesg = SUCCESS_CODE, SUCCESS_MESG
    }
    c.JSON(http.StatusOK, rep)
  }
}

func GetTemplate(c *gin.Context) {
  req, rep := new(GetTemplateReq), new(GetTemplateRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "template.get",
      "reason": err.Error(),
    }).Error(GIN_PARSE_REQUEST_LOG)
    rep.Code = GIN_PARSE_REQUEST_CODE
    rep.Mesg = err.Error()
  } else if data := models.TemplateCache.Get(req.Name); data != nil {
    log.WithFields(log.Fields{
      "action": "template.get",
      "req": req,
      "reason": err.Error(),
    }).Error(CACHE_GET_DATA_LOG)
    rep.Code = CACHE_GET_DATA_CODE
    rep.Mesg = "template not found"
  } else if templ, ok := data.(*models.Template); !ok {
    log.WithFields(log.Fields{
      "action": "template.get",
      "req": req,
      "reason": "type assert failed",
    }).Error(CACHE_GET_TYPE_LOG)
    rep.Code = CACHE_GET_TYPE_CODE
    rep.Mesg = "template not found"
  } else {
    rep.Code, rep.Mesg = SUCCESS_CODE, SUCCESS_MESG
    rep.Name, rep.Config = templ.Name, templ.Config
  }
  c.JSON(http.StatusOK, rep)
}

func DelTemplate(c *gin.Context) {
  req, rep := new(DelTemplateReq), new(DelTemplateRep)
  if err := util.ReadJSON(c.Request, req); err != nil {
    log.WithFields(log.Fields{
      "action": "template.del",
      "reason": err.Error(),
    }).Error(GIN_PARSE_REQUEST_LOG)
    rep.Code = GIN_PARSE_REQUEST_CODE
    rep.Mesg = err.Error()
  } else if err := models.TemplateCache.Delete(req.Name); err != nil {
    log.WithFields(log.Fields{
      "action": "template.del",
      "reason": err.Error(),
    }).Error(CACHE_DEL_DATA_LOG)
    rep.Code = CACHE_DEL_DATA_CODE
    rep.Mesg = err.Error()
  } else {
    rep.Code, rep.Mesg = SUCCESS_CODE, SUCCESS_MESG
  }
  c.JSON(http.StatusOK, rep)
}
