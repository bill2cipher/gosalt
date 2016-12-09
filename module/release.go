/*
 * This module is used to make release packages for gosalt
**/
package module

import (
  "github.com/spf13/viper"
  "github.com/jellybean4/gosalt/util"
  log "github.com/Sirupsen/logrus"
)

/**
 * Args passed to release script include:
 * 1. code dir
 * 2. version
 * 3. release dir
 * 4. server types
 */
func Release(version string, types ...string) *util.ExecResult {
  releaseDir := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.RELEASE_DIR)
  if err := util.CheckDir(releaseDir + "/" + version, true); err != nil {
    log.WithFields(log.Fields{
      "dir": releaseDir + "/" + version,
      "reason": err.Error(),
    }).Error(util.DIR_CREAE_LOG)
  }
  args := []string{viper.GetString(util.CODE_DIR), version, releaseDir}
  args = append(args, types...)
  releaseScript := viper.GetString(util.ROOT_DIR) + "/" +
      viper.GetString(util.RELEASE_SCRIPT)
  log.WithFields(log.Fields{
    "script": releaseScript,
    "args": args,
  }).Debug("starting execute script")
  return util.ExecScript(releaseScript, args...)
}
