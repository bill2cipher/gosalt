package util

import (
  "github.com/spf13/viper"
  "fmt"
)

func GetConfig(name string) string {
  switch name {
  case MINION_DIR:
    return fmt.Sprintf("%s/%s", viper.GetString(ROOT_DIR), viper.GetString(MINION_DIR))
  case CONFIG_TEMPL_DIR:
    return fmt.Sprintf("%s/%s", viper.GetString(ROOT_DIR), viper.GetString(CONFIG_TEMPL_DIR))
  default:
    return viper.GetString(name)
  }
}
