package vault

import (
  "testing"
)

func TestEndpointExists(t *testing.T) {

  t.Run("Endpoint exists", func(t *testing.T) {
    got := EndpointExists("admin")
    want := true

    if got != want {
      t.Errorf("got %t want %t", got, want)
    }
  })

  t.Run("Endpoint doesn't exist", func(t *testing.T) {
    got := EndpointExists("nope")
    want := false

    if got != want {
      t.Errorf("got %t want %t", got, want)
    }
  })

}

func TestVaultApi(t *testing.T) {

  validURL := "http://vault.local.seanhome.xyz/"
  validAPP := "admin"
  /*
  bogusURL := "http://vault.fake.seanhome.xyz/"
  bogusAPP := "fake"
  */

  t.Run("Returns 200 status code", func(t *testing.T) {
    got := VaultApi(validURL, validAPP)
    want := 200

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  /*
  t.Run("Returns 200 status code", func(t *testing.T) {
    got := VaultApi(bogusURL, bogusAPP)
    want := 200

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })
  */

}
