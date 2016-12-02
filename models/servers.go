package models

import (
  "github.com/jellybean4/gosalt/db"
  "github.com/jellybean4/gosalt/util"
  log "github.com/Sirupsen/logrus"
  "github.com/spf13/viper"
)

var (
  ServerCache *Cache
)

func init() {
  ServerCache = NewServerCache()
}

type (
  Server struct {
    Name     string            `json:"name"`
    IPAddr   string            `json:"ipaddr"`
    SshPort  int               `json:"sshport"`
    UserName string            `json:"username"`
    Passwd   string            `json:"passwd"`
    Property map[string]string `json:"property"`
  }

  ServerHandler struct {
    *DefaultHandler
  }
)

func (h *ServerHandler) Unmarshal(data []byte) (interface{}, error) {
  v := new(Server)
  if err := util.Unmarshal(data, v); err != nil {
    log.WithFields(log.Fields{
      "data": data,
      "reason": err.Error(),
    }).Error("unmarshal data failed")
    return nil, err
  } else {
    return v, nil
  }
}

func (h *ServerHandler) Key(value interface{}) (string, bool) {
  if s, ok := value.(*Server); ok {
    log.WithFields(log.Fields{
      "value": value,
      "type": "server",
    }).Error("type assert failed")
    return s.Name, true
  } else {
    return "", false
  }
}

func NewServerCache() *Cache {
  dftHandler := &DefaultHandler{table : db.SERVER_TABLE}
  svrHandler := &ServerHandler{DefaultHandler: dftHandler}
  c := NewCache(string(db.SERVER_TABLE), svrHandler, true, false)
  return c
}

// InitConfig is used to init the server's salt minion config
func (svr *Server) InitConfig(args ...string) *util.ExecResult {
  initScript := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.INIT_SCRIPT)
  return util.ExecScript(initScript, args)
}

// SyncConfig is used to sync config info into minion
func (svr *Server) SyncConfig(args ...string) *util.ExecResult {
  syncScript := viper.GetString(util.ROOT_DIR) + "/" + viper.GetString(util.SYNC_SCRIPT)
  return util.ExecScript(syncScript, args)
}
