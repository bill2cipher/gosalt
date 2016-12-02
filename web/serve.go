package web

import (
  "github.com/gin-gonic/gin"
  "fmt"
)

import (
  . "github.com/jellybean4/gosalt/web/controller"
  "github.com/jellybean4/gosalt/util"
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/release"
  "github.com/jellybean4/gosalt/deploy"
)

func Serve(port int) {
  router := gin.Default()
  webDir := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.WEB_DIR)
  static(webDir, router)
  routes(webDir, router)

  lport := fmt.Sprintf(":%d", port)
  router.Run(lport)
}

func static(webDir string, router *gin.Engine) {
  router.Static("/pages", webDir + "/pages")
  router.Static("/data", webDir + "/data")
  router.Static("/dist", webDir + "/dist")
  router.Static("/js", webDir + "/js")
  router.Static("/less", webDir + "/less")
  router.Static("/vendor", webDir + "/vendor")
}

func routes(webDir string, router *gin.Engine) {
  router.StaticFile("index.html", webDir + "/index.html")
  release := router.Group("/new")
  {
    release.POST("/add", Release)
  }

  deploy := router.Group("/deploy")
  {
    deploy.POST("/do", Deploy)
  }

  config := router.Group("/config")
  {
    config.POST("/init", InitServer)
    config.POST("/set", SetServer)
    config.POST("/del", DelServer)
    config.POST("/get", GetServer)
  }

  template := router.Group("/template")
  {
    template.POST("/set", SetTemplate)
    template.POST("/get", GetTemplate)
    template.POST("/del", DelTemplate)
  }
}
