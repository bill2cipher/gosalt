// Copyright © 2016 jellybean4
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "fmt"
  log "github.com/Sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "os"
)

import (
  "github.com/jellybean4/gosalt/util"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
  Use:   "gosalt",
  Short: "gosalt is a tool used to deploy and config servers with saltstack",
  Long: `saltstack is a great tool for server manage, but to make it easy for
use, here's gosalt`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/gosalt/gosalt.yml)")
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
  viper.SetDefault(util.DB_FILE, "/var/gosalt/gosaltdb")
  viper.SetDefault(util.WEB_DIR, "/etc/gosalt/web")

  viper.SetDefault(util.RELEASE_SCRIPT, "/etc/gosalt/script/release.sh")
  viper.SetDefault(util.INIT_SCRIPT, "/etc/gosalt/script/init.sh")
  viper.SetDefault(util.SYNC_SCRIPT, "/etc/gosalt/script/sync.sh")
  viper.SetDefault(util.DEPLOY_SCRIPT, "/etc/gosalt/script/deploy.sh")
}

