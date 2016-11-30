package util

import (
  "bytes"
  "os/exec"
  "encoding/json"
)

import (
  log "github.com/Sirupsen/logrus"
)

// Diec is used to indicate the exiting of gosalt server
var Diec chan *struct{}


func init() {
  Diec = make(chan *struct{})
}

func Marshal(v interface{}) ([]byte, error) {
  return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
  return json.Unmarshal(data, v)
}

func ExecScript(script string, args ...string) {
  cmd := exec.Command(script, args...)
  var errBuff, outBuff bytes.Buffer
  cmd.Stderr, cmd.Stdout = &errBuff, &outBuff
  if err := cmd.Run(); err != nil {
    log.WithFields(log.Fields{
      "script": script,
      "errout": errBuff.String(),
      "stdout": outBuff.String(),
      "args": args,
      "reason": err.Error(),
    }).Error("execute script failed")
  } else {
    log.WithFields(log.Fields{
      "script": script,
      "errout": errBuff.String(),
      "stdout": outBuff.String(),
      "args": args,
    }).Info("execute script success")
  }
}
