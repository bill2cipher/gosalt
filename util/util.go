package util

import (
  "encoding/json"
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
