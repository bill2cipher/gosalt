package saltstack

import "encoding/json"
import (
	log "github.com/Sirupsen/logrus"
)

/**
 * Use saltstack modules wheel to read master config data,
 * manage template files or manage pillar state files.
 */

func SaltWheel(f string, low map[string]interface{}) (*json.RawMessage, error) {
	low["client"] = "wheel"
	low["fun"] = f

	if raw, err := runLowStat(low); err != nil {
		return nil, err
	} else {
		log.WithFields(log.Fields{
			"raw": raw,
			"low": low,
		}).Info("wheel accept key success")
		return raw, nil
	}
}

func AcceptMinion(name string) error {
	low := make(map[string]interface{})
	low["match"] = []string{name}

	if _, err := SaltWheel("key.accept", low); err != nil {
		return err
	} else {
		return nil
	}
}
