package config

import (
  "testing"
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
