package test

import (
"encoding/json"
"testing"
)

import (
. "github.com/jellybean4/gosalt/mesg"
)

func TestRelease(t *testing.T) {
  startServer(t)

  reqStruct := &ReleaseReq{
    Version: "trunk",
    Types:   []string{"all"},
  }

  data, repStruct := request(t, reqStruct, "release/do"), &ReleaseRep{}
  if err := json.Unmarshal(data, repStruct); err != nil {
    t.Errorf("parse release response failed, reason %s", err.Error())
  }
}
