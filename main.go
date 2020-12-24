package main

import (
    f "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"

    "curator/pkg/config"
    //"curator/pkg/vault"
    //"curator/pkg/spinnaker"
)

type Conf struct {
    Apps []map[string]string `yaml:"rails_apps"`
}

/*
type app struct {
  App map[string]string
}
*/

func (c *Conf) GetConf() *Conf {

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c Conf
    c.GetConf()

    for _, app := range c.Apps {
      f.Printf("Name: %s\n", app["name"])
      f.Printf("Repo: %s\n", app["repo"])
      //vault.Endpoint(app["name"], app["repo"])
      //spinnaker.Pipeline(app["name"], app["repo"])
    }
}
