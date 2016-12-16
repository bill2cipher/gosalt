package util

import (
  "bytes"
  "encoding/json"
  "os/exec"
  "errors"
  "net/http"
  "os"
  "fmt"
)


// Diec is used to indicate the exiting of gosalt server
var (
  Diec chan *struct{}
)

type ExecResult struct {
  Error  error
  Script string
  Args   []string
  Stdout *bytes.Buffer
  Stderr *bytes.Buffer
}

func initUtil() {
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

func OpenFile(dir, name string, create bool) (*os.File, error) {
  filename := dir + "/" + name
  if err := CheckDir(dir, create); err != nil {
    return nil, err
  }
  if info, err := os.Stat(filename); err == nil && info.IsDir() {
    mesg := fmt.Sprintf("%s is directory", filename)
    return nil, errors.New(mesg)
  } else if err != nil && !create {
    return nil, err
  }

  flags := os.O_TRUNC | os.O_CREATE | os.O_RDWR
  if f, err := os.OpenFile(filename, flags, MODE); err != nil {
    return nil, err
  } else {
    return f, nil
  }
}

// CheckDir check the existence of given dir, if not, create it.
func CheckDir(dir string, create bool) error {
  if info, err := os.Stat(dir); err != nil {
    if create {
      if err := os.MkdirAll(dir, MODE); err != nil {
        return err
      }
      return nil
    }
    return err
  } else if info.IsDir() {
    return nil
  } else {
    mesg := fmt.Sprintf("%s not directory", dir)
    return errors.New(mesg)
  }
}

func Recover() error {
  if err := recover(); err != nil {
    mesg := fmt.Sprintf("%v")
    return errors.New(mesg)
  }
  return nil
}
