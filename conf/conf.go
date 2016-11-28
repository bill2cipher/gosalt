package conf

import (
	"github.com/jinzhu/configor"
)

// GoSaltConf contains all configuration for gosalt
type GoSaltConf struct {
	CodeDir       string `default:"/root/p4/Programe/trunc/Server"`
	ReleaseScript string `default:"/etc/gosalt/release.sh"`
	ReleaseDir    string `default:"/var/gosalt/release"`
	DBFile        string `default:"/etc/gosalt/gosaltdb"`
}

// Conf is the global gosalt conf variable of type GoSaltConf
var (
	Conf   *GoSaltConf
	Tables = []string{"server"}
)

// Init action for gosalt configuration
func initConf() {
	Conf = new(GoSaltConf)
}

func Load(configFile string) {
	configor.Load(Conf, configFile)
}
