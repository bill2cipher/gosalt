package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
)

import (
	"errors"
	"reflect"
)

var (
	ServerCache *Cache
)

func initServer() {
	if c, err := NewServerCache(); err != nil {
		log.WithFields(log.Fields{
			"table":  db.SERVER_TABLE,
			"reason": err.Error(),
		}).Fatal(util.CACHE_INIT_DATA_LOG)
	} else {
		ServerCache = c
	}
}

type (
	Server struct {
		Name     string `json:"name"     bson:"_id"`
		IPAddr   string `json:"ipaddr"   bson:"ipaddr"`
		SshPort  int    `json:"sshport"  bson:"sshport"`
		UserName string `json:"username" bson:"username"`
		Passwd   string `json:"passwd"   bson:"passwd"`
		Env      string `json:"env"      bson:"env"`
		Version  string `json:"version" bson:"version"`
	}

	ServerHandler struct {
		*DefaultHandler
	}
)

func (h *ServerHandler) SetKey(v interface{}, key string) error {
	return errors.New("server model not support")
}

func (h *ServerHandler) Key(value interface{}) (string, bool) {
	if s, ok := value.(*Server); !ok {
		log.WithFields(log.Fields{
			"value": value,
			"type":  "server",
		}).Error(util.TYPE_ASSERT_LOG)
		return s.Name, true
	} else {
		return "", false
	}
}

func NewServerCache() (*Cache, error) {
	dftHandler := &DefaultHandler{table: db.SERVER_TABLE, t: reflect.TypeOf(Server{})}
	svrHandler := &ServerHandler{DefaultHandler: dftHandler}
	return NewCache(string(db.SERVER_TABLE), svrHandler, true, false)
}
