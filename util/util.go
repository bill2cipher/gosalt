package util

import (
  "bytes"
  "encoding/json"
  "os/exec"
  "errors"
  "net/http"
)

// Diec is used to indicate the exiting of gosalt server
var Diec chan *struct{}

type ExecResult struct {
  Error  error
  Script string
  Args   []string
  Stdout *bytes.Buffer
  Stderr *bytes.Buffer
}

func init() {
  Diec = make(chan *struct{})
}

func Marshal(v interface{}) ([]byte, error) {
  return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
  return json.Unmarshal(data, v)
}

func ExecScript(script string, args ...string) *ExecResult {
  cmd := exec.Command(script, args...)
  errBuff, outBuff := new(bytes.Buffer), new(bytes.Buffer)
  cmd.Stderr, cmd.Stdout = errBuff, outBuff
  result := &ExecResult{
    Error:  nil,
    Script: script,
    Args:   args,
    Stdout: outBuff,
    Stderr: errBuff,
  }
  if err := cmd.Run(); err != nil {
    result.Error = err
  }
  return result
}

func ReadJSON(request *http.Request, v interface{}) error {
  if data, err := ReadContent(request); err != nil {
    return err
  } else {
    return json.Unmarshal(data, v)
  }
}

func ReadContent(request *http.Request) ([]byte, error) {
  data := make([]byte, request.ContentLength)
  if n, err := request.Body.Read(data); n == 0 && err != nil {
    return nil, err
  } else if n != int(request.ContentLength) {
    return nil, errors.New("content length not match")
  } else {
    return data, nil
  }
}
