package test

import "fmt"
import "encoding/json"
import "testing"

func TestGo(t *testing.T) {
	var A struct {
		A int
		B map[string]string
	}
	A.A = 12
	A.B = map[string]string{"good": "bad", "think": "follow"}
	B, _ := json.Marshal(A)
	fmt.Println("%v", B)
}
