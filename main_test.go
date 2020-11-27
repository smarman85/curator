package main

import (
  "testing"
  "log"
  "os"
)

func init() {
  filename := "config.yaml"
  _, err := os.Stat(filename)
  if err != nil {
    log.Fatalf("Could not find %s in current directory.", filename)
  }
}

func TestGetConf(t *testing.T) {
  var c Conf
  first := c.GetConf().Apps[0]["name"]
  want:= "authtest"

  if first != want {
    t.Errorf("got %s want %s", first, want)
  }
}
