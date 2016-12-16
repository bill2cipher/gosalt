package saltstack

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/mesg"
	"github.com/jellybean4/gosalt/util"
	"github.com/parnurzeal/gorequest"
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	mu          sync.RWMutex
	globalToken = new(mesg.SaltLoginRep)
	REQ_TIMEOUT = 5 * time.Second
)

type SaltLowState map[string]interface{}

func saltToken(force bool) (string, error) {
	mu.RLock()
	now, expire, token := time.Now().Unix(), int64(globalToken.Expire), globalToken.Token
	mu.RUnlock()

	if force || now >= expire {
		mu.Lock()
		defer mu.Unlock()

		if err := saltLogin(); err != nil {
			return "", err
		} else {
			return globalToken.Token, nil
		}
	} else {
		return token, nil
	}
}

func saltLogin() error {
	url := fmt.Sprintf("http://%s:%s/login", util.GetConfig(util.MASTER),
		util.GetConfig(util.NETAPI_PORT))

	loginReq := &mesg.SaltLoginReq{
		User:  util.GetConfig(util.MASTER_USER),
		Pass:  util.GetConfig(util.MASTER_PASS),
		EAuth: "pam",
	}
	var loginRep []*mesg.SaltLoginRep
	result := make(map[string]*json.RawMessage)
	resp, body, errs := gorequest.New().Post(url).
		Set("Accept", "application/json").
		Send(loginReq).
		Timeout(REQ_TIMEOUT).
		EndBytes()

	if len(errs) != 0 {
		log.WithFields(log.Fields{
			"user":   util.GetConfig(util.MASTER_USER),
			"pass":   util.GetConfig(util.MASTER_PASS),
			"reason": errs,
		}).Error(util.SALT_LOGIN_LOG)
		return errs[0]
	} else if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"user":   util.GetConfig(util.MASTER_USER),
			"pass":   util.GetConfig(util.MASTER_PASS),
			"status": resp.Status,
		}).Error(util.SALT_LOGIN_LOG)
		return errors.New("login resp status not ok")
	} else if err := json.Unmarshal(body, &result); err != nil {
		log.WithFields(log.Fields{
			"data":   body,
			"reason": err.Error(),
		}).Error(util.JSON_UNMARSHAL_LOG)
		return errors.New(util.JSON_UNMARSHAL_LOG)
	} else if ret, ok := result["return"]; !ok {
		log.WithFields(log.Fields{
			"data":   result,
			"reason": "return format error",
		}).Error(util.SALT_LOGIN_LOG)
		return errors.New("token not exist")
	} else if err := json.Unmarshal(*ret, &loginRep); err != nil {
		log.WithFields(log.Fields{
			"return": ret,
			"reason": err.Error(),
		}).Error(util.SALT_LOGIN_LOG)
		return err
	} else if len(loginRep) <= 0 {
		log.WithFields(log.Fields{
			"return": ret,
			"reason": "token not exist",
		}).Error(util.SALT_LOGIN_LOG)
		return errors.New("token not exist")
	} else {
		globalToken = loginRep[0]
		return nil
	}
}

func runLowStat(low SaltLowState) (*json.RawMessage, error) {
	url := fmt.Sprintf("http://%s:%s", util.GetConfig(util.MASTER),
		util.GetConfig(util.NETAPI_PORT))

	result := make(map[string]*json.RawMessage)
	retry := false
	token, err := saltToken(false)
	if err != nil {
		return nil, err
	}

RETRY:
	if retry {
		token, err = saltToken(true)
		if err != nil {
			return nil, err
		}
	}
	resp, body, errs := gorequest.New().
		Timeout(REQ_TIMEOUT).
		Post(url).
		Set("Accept", "application/json").
		Set("X-Auth-Token", token).
		Send(low).
		EndBytes()

	if len(errs) != 0 {
		log.WithFields(log.Fields{
			"lowstate": low,
			"reason":   errs,
		}).Error(util.SALT_EXEC_LOG)
		return nil, errs[0]
	} else if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"lowstate": low,
			"status":   resp.Status,
		}).Error(util.SALT_EXEC_LOG)
		if retry {
			return nil, errors.New("exec resp status not ok")
		} else {
			retry = true
			goto RETRY
		}
	} else if err := json.Unmarshal(body, &result); err != nil {
		log.WithFields(log.Fields{
			"lowstate": low,
			"data":     body,
			"reason":   err.Error(),
		}).Error(util.JSON_UNMARSHAL_LOG)
		return nil, errors.New(util.JSON_UNMARSHAL_LOG)
	} else if ret, ok := result["return"]; !ok {
		log.WithFields(log.Fields{
			"lowstate": low,
			"data":     result,
			"reason":   "return format error",
		}).Error(util.SALT_EXEC_LOG)
		return nil, errors.New("return format error")
	} else {
		return ret, nil
	}
}

func SaltExec(target string, f string, arg ...string) (*json.RawMessage, error) {
	low := SaltLowState{"client": "local", "tgt": target, "fun": f}
	if len(arg) != 0 {
		low["arg"] = arg
	}
	if raw, err := runLowStat(low); err != nil {
		return nil, err
	} else {
		return raw, nil
	}
}
