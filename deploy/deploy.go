package deploy

import (
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
  "github.com/jellybean4/gosalt/models"
)

func Deploy(version string, servers string) *util.ExecResult {
  args := []string{viper.GetString(util.RELEASE_DIR), version, servers}
  deployScript := viper.GetString(util.ROOT_DIR) + "/" +
      viper.GetString(util.DEPLOY_SCRIPT)

  return util.ExecScript(deployScript, args...)
}

