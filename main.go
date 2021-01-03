package main

import (
    f "fmt"
    //"gopkg.in/yaml.v2"
    //"io/ioutil"
    "log"
    "os"

    "curator/pkg/config"
    //"curator/pkg/vault"
    //"curator/pkg/spinnaker"
)

var fileName string = os.Args[1]

type Conf struct {
    Apps []map[string]string `yaml:"rails_apps"`
}

/*
type app struct {
  App map[string]string
}
*/

/*
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
*/

func main() {
  rawConfig, err := config.GetConfig(fileName)
  if err != nil {
    log.Printf("%v", err)
  }

  data, _ := config.DecodeConfig(rawConfig)
  f.Println(data.Apps[0]["name"])
  //for _, app := range data.Apps {
  //  f.Println(app["name"])
  //}

  /*
    var c Conf
    c.GetConf()

    for _, app := range c.Apps {
      f.Printf("Name: %s\n", app["name"])
      f.Printf("Repo: %s\n", app["repo"])
      //vault.Endpoint(app["name"], app["repo"])
      //spinnaker.Pipeline(app["name"], app["repo"])
    }
  */
}
