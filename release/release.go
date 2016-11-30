/*
 * This module is used to make release packages for gosalt
**/
package release

import (
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
)

/**
 * Args passed to release script include:
 * 1. code dir
 * 2. version
 * 3. release dir
 * 4. server types
 */
func Release(version string, types ...string) {
  args := []string{viper.GetString(util.CODE_DIR), version,
      viper.GetString(util.RELEASE_DIR)}
  args = append(args, types...)
  releaseScript := viper.GetString(util.ROOT_DIR) + "/" +
      viper.GetString(util.RELEASE_SCRIPT)

  util.ExecScript(releaseScript, args...)
}
