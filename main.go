package main

import (
    f "fmt"
    //"gopkg.in/yaml.v2"
    //"io/ioutil"
    "log"
    "os"

    "curator/pkg/config"
    "curator/pkg/vault"
    //"curator/pkg/spinnaker"
)

var fileName string = os.Args[1]

type Conf struct {
    Apps []map[string]string `yaml:"rails_apps"`
}

func main() {
  rawConfig, err := config.GetConfig(fileName)
  if err != nil {
    log.Printf("%v", err)
  }

  data, _ := config.DecodeConfig(rawConfig)
  f.Println(data.Apps[0]["name"])
}
