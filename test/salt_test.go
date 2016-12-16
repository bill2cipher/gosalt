package test

import (
  "testing"
  "github.com/jellybean4/gosalt/saltstack"
  "time"
)

func TestSalt(t *testing.T) {
  if rslt, err := saltstack.SaltExec("*", "cmd.run", "cmd=ls -lsh", "cwd=/tmp"); err != nil {
    t.Error(err.Error())
  } else {
    t.Logf("salt exec result %s", *rslt)
  }
}

func TestWheelAcceptKey(t *testing.T) {
  if err := saltstack.AcceptMinion("test"); err != nil {
    t.Error(err.Error())
  }
}

func TestEvent(t *testing.T) {
  time.Sleep(5 * time.Minute)
}
