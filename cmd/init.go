package cmd

import (
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/order_init"
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
)

func init() {
  order.RegisteFunc(util.CMD_INIT, initCmd)
  order.RegisteFunc(util.CONFIG_INIT, initConfig, util.CMD_INIT, util.LOG_INIT)
}

func initCmd() {
  initRoot()
  initRelease()
  initStart()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  setDefault() // set default config for gosalt

  if cfgFile != "" {
    // enable ability to specify config file via flag
    viper.SetConfigFile(cfgFile)
  }

  viper.SetConfigName("gosalt")      // name of config file (without extension)
  viper.AddConfigPath("/etc/gosalt") // adding home directory as first search path
  viper.AutomaticEnv()               // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err != nil {
    log.WithFields(log.Fields{
      "config": viper.ConfigFileUsed(),
      "reason": err.Error(),
    }).Fatal("parse config file failed")
  } else {
    log.WithFields(log.Fields{
      "config": viper.ConfigFileUsed(),
    }).Info("parse config file success")
  }
}

func setDefault() {
  viper.SetDefault(util.ROOT_DIR, "/etc/gosalt/")
  viper.SetDefault(util.CODE_DIR, "/root/p4/Programe/trunc/Server/")
  viper.SetDefault(util.RELEASE_DIR, "/var/gosalt/release/")
  viper.SetDefault(util.WEB_DIR, "/etc/gosalt/web")

  viper.SetDefault(util.RELEASE_SCRIPT, "/etc/gosalt/script/release.sh")
  viper.SetDefault(util.INIT_SCRIPT, "/etc/gosalt/script/init.sh")
  viper.SetDefault(util.SYNC_SCRIPT, "/etc/gosalt/script/sync.sh")
  viper.SetDefault(util.DEPLOY_SCRIPT, "/etc/gosalt/script/deploy.sh")
}
