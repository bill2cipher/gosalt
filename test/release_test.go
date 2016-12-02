package test

import (
  "encoding/json"
  "testing"
)

import (
  . "github.com/jellybean4/gosalt/mesg"
)

func TestRelease(t *testing.T) {
  initServer(t)

  reqStruct := &ReleaseReq{
    Version: "trunk",
    Types:   []string{"all"},
  }

  data, repStruct := request(t, reqStruct, "action/release"), &ReleaseRep{}
  if err := json.Unmarshal(data, repStruct); err != nil {
    t.Errorf("parse release response failed, reason %s", err.Error())
  }
}
