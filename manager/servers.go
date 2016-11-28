package manager

import (
	"fmt"
	"github.com/jellybean4/gosalt/db"
	"github.com/jellybean4/gosalt/util"
)

import (
	log "github.com/Sirupsen/logrus"
)

var Servers map[string]Server

const (
	SVR_PROP_TYPE = "type"
)

type Server struct {
	Name     string            `json:"name"`
	IPAddr   string            `json:"ipaddr"`
	SshPort  int               `json:"sshport"`
	UserName string            `json:"username"`
	Passwd   string            `json:"passwd"`
	Property map[string]string `json:"property"`
}

func init() {
	if datas, err := db.All(db.SERVER_TABLE); err != nil {
		mesg := fmt.Sprintf("load server data failed, reason %s", err.Error())
		panic(mesg)
	} else {
		for data := range datas {
			svr := new(Server)
			if err := util.Unmarshal(data, svr); err != nil {
				mesg := fmt.Sprintf("unmarshal server data %v failed, reason %s", data, err.Error())
				panic(mesg)
			} else {
				Servers[svr.Name] = svr
			}
		}
	}
}

func AddServer(svrData string) (*Server, error) {
  svr := new(Server)
  if err := util.Unmarshal(svrData, svr); err != nil {
    log.WithFields(log.Fields{
      "data": svrData,
      "reason": err.Error(),
    }).Error("unmarshal server data failed")
    return nil, err
  } else if db.Set(db.SERVER_TABLE, []byte(name), svrData); err != nil {
		return nil, err
	} else {
		Servers[svr.Name] = svr
		return svr, nil
	}
}

func DelServer(name string) error {
	if err := db.Unset(db.SERVER_TABLE, []byte(name)); err != nil {
		return err
	} else {
		delete(Servers, name)
		return nil
	}
}

// InitConfig is used to init the server's salt minion config
func (svr *Server) InitConfig() {
  cmd := ""
}

// SyncConfig is used to sync config info into minion
func (svr *Server) SyncConfig() {

}
