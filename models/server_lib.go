package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/util"
	"path"
	"text/template"
)

const (
	MINION_TEMPLATE = `master: {{.master_ip}}
publish_port: {{.master_port}}
user: {{.minion_user}}
root_dir: {{.minion_root}}
id: {{.minion_name}}
environment: {{.minion_env}}
grains:
  - version: {{.minion_version}}
`
)

var (
	serverParser *template.Template
)

func initServerLib() {
	if templ, err := template.New("server").Parse(MINION_TEMPLATE); err != nil {
		log.WithFields(log.Fields{
			"template": "server",
			"reason":   err.Error(),
		}).Fatal(util.TEMPL_PARSE_LOG)
	} else {
		serverParser = templ
	}
}

// SetServer generate saltstack config for minion server and restart
// saltstack minion.
func SetServer(svr *Server) error {
	minionDir := util.GetConfig(util.MINION_DIR)
	writer, err := util.OpenFile(minionDir, svr.Name, true)
	if err != nil {
		log.WithFields(log.Fields{
			"file":   path.Join(minionDir, svr.Name),
			"reason": err.Error(),
		}).Error(util.FILE_OPEN_LOG)
		return err
	}

	args := make(map[string]string)
	args["master_ip"] = util.GetConfig(util.MASTER)
	args["master_port"] = util.GetConfig(util.MASTER_PORT)
	args["minion_user"] = util.GetConfig(util.MINION_USER)
	args["minion_root"] = util.GetConfig(util.MINION_ROOT)
	args["minion_name"] = svr.Name
	args["minion_env"] = svr.Env
	args["minion_version"] = svr.Version

	if err := serverParser.Execute(writer, args); err != nil {
		log.WithFields(log.Fields{
			"args":   args,
			"reason": err.Error(),
		}).Error(util.TEMPL_EXEC_LOG)
		return err
	} else if err := writer.Close(); err != nil {
		log.WithFields(log.Fields{
			"file":   path.Join(minionDir, svr.Name),
			"reason": err.Error(),
		}).Error(util.FILE_CLOSE_LOG)
		return err
	} else if err := ServerCache.Set(svr.Name, svr); err != nil {
		return err
	} else {
		log.WithFields(log.Fields{
			"file": path.Join(minionDir, svr.Name),
			"args": args,
		}).Info("template execute success")
		return nil
	}
}

/**
 * SyncSaltConfig use salt ssh to sync salt minion config into minion
 */
func SyncSaltConfig(svr *Server) error {
	return nil
}

func SyncServerConfig() error {
	return nil
}

func GenerateRSAKey() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		log.WithFields(log.Fields{
			"reason": err.Error(),
		}).Error(util.RSA_GEN_LOG)
		return "", "", err
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.WithFields(log.Fields{
			"reason": err.Error(),
			"key":    publicKey,
		}).Error(util.RSA_MARSHAL_LOG)
		return "", "", err
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&publicKeyBlock))

	return privateKeyPem, publicKeyPem, nil
}
