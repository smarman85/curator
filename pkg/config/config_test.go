package config

import (
  "testing"
  //"reflect"
)

func assertString(t *testing.T, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func assertError(t *testing.T, got, want error) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func TestGetConfig(t *testing.T) {

  t.Run("errors on missing file", func(t *testing.T) {
    _, got := GetConfig("missing.yaml")

    assertError(t, got, MissingFile)
  })

}

func TestDecodeConfig(t *testing.T) {

  //yamlData := []byte(`{"rails_apps":[{"name":"sampleapp","repo":"sample-app"}]}`)


  t.Run("test returns content", func(t *testing.T) {
    data, _ := GetConfig("test.yaml")
    got, _ := DecodeConfig(data)
    app := got.Apps[0]["name"]
    want := "sampleapp"

    if app != want {
      t.Errorf("got %q want %q", got, want)
    }

  })
}
