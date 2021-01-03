package config

import (
  //f "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

const (
  MissingFile = ConfigErr("could not find the config file")
  UnmarshalError = ConfigErr("could not unmarshal yaml data")
)

type ConfigErr string

func (c ConfigErr) Error() string {
  return string(c)
}

type Config struct {
  Apps []map[string]string `yaml:"rails_apps"`
}

func DecodeConfig(yamlData []byte) (Config, error) {
  var c Config
  err := yaml.Unmarshal(yamlData, &c)
  if err != nil {
    return c, err
  }
  return c, nil
}

func GetConfig(fileName string) ([]byte, error) {
  yamlFile, err := ioutil.ReadFile(fileName)
  if err != nil {
    return nil, MissingFile
    //return "", "missing file"
  }
  //f.Printf("%T\n", yamlFile)
  //f.Println(string(yamlFile))
  return yamlFile, nil
}
