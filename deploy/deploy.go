package deploy

import (
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
)

func Deploy(version string, servers string) {
  args := []string{viper.GetString(util.RELEASE_DIR), version, servers}
  deployScript := viper.GetString(util.ROOT_DIR) + "/" +
      viper.GetString(util.DEPLOY_SCRIPT)

  util.ExecScript(deployScript, args...)
}
