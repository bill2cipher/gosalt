package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
)

var (
	ServerCache *Cache
)

func initServer() {
	if c, err := NewServerCache(); err != nil {
    log.WithFields(log.Fields{
      "table": db.SERVER_TABLE,
      "reason": err.Error(),
    }).Fatal(util.CACHE_INIT_DATA_LOG)
  } else {
    ServerCache = c
  }
}

type (
	Server struct {
		Name     string `json:"name"`
		IPAddr   string `json:"ipaddr"`
		SshPort  int    `json:"sshport"`
		UserName string `json:"username"`
		Passwd   string `json:"passwd"`
		Env      string `json:"env"`
	}

	ServerHandler struct {
		*DefaultHandler
	}
)

func (h *ServerHandler) Key(value interface{}) (string, bool) {
	if s, ok := value.(*Server); ok {
		log.WithFields(log.Fields{
			"value": value,
			"type":  "server",
		}).Error("type assert failed")
		return s.Name, true
	} else {
		return "", false
	}
}

func NewServerCache() (*Cache, error) {
	dftHandler := &DefaultHandler{table: db.SERVER_TABLE}
	svrHandler := &ServerHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.SERVER_TABLE), svrHandler, true, false)
}
