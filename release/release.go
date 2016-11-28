/*
 * This module is used to make release packages for
**/
package release

import (
	"bytes"
	"os/exec"
)

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/jellybean4/gosalt/conf"
)

func Release(args ...string) {
	ReleaseScript := Conf.RootDir + "/" + Conf.ReleaseScript
	cmd := exec.Command(ReleaseScript, args...)
	var errBuff, outBuff bytes.Buffer
	cmd.Stderr, cmd.Stdout = &errBuff, &outBuff
	if err := cmd.Run(); err != nil {
		log.Printf("run release script failed, reason %s", err.Error())
	} else {
		log.Printf("run release script success, output %s", outBuff.String())
	}
}
