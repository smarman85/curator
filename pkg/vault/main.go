package vault

import (
  "io/ioutil"
  "bytes"
  "net/http"
  "encoding/json"
  f "fmt"
  //"log"
  "os"
)

var VltPass string = os.Getenv("VAULT_TOKEN")

type sample struct {
  Data string `json:"env"`
}

var vltUrl string = "http://vault.local.seanhome.xyz/"

func check(msg string, e error) {
  if e != nil {
    f.Fprintf(os.Stderr, msg + "\n", e)
 }
}

func vaultApi(url, app string) int {
  request, err := http.NewRequest("GET", url + "v1/secret/data/test/" + app + "/env", nil)
  check("Listing vault urls", err)
  request.Header.Set("X-Vault-Token", VltPass)

  resp, err := http.DefaultClient.Do(request)
  check("making request to vault", err)
  return resp.StatusCode
  /*
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  check("reading response from vault", err)
  return body
  */
}

func createVaultEpt(url, app string) []byte {
  jData, _ := json.Marshal(map[string]map[string]string{
    "data": {"env":"test"},
  })
  request, err := http.NewRequest("POST", url + "v1/secret/data/test/" + app + "/env", bytes.NewBuffer(jData))
  check("Listing vault urls", err)
  request.Header.Set("X-Vault-Token", VltPass)
  request.Header.Add("Content-Type", "application/json")


  resp, err := http.DefaultClient.Do(request)
  check("making request to vault", err)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  check("reading response from vault", err)
  return body
}

func endPointExists(app string) bool {
  rcode := vaultApi(vltUrl, app)
  exists := false
  if rcode >= 200 && rcode <= 299 {
    exists = true
  }
  return exists
}

func Endpoint(app string) {
  if !endPointExists(app) {
    f.Println("Creating Endpoint for: " + app)
    resp := createVaultEpt(vltUrl, app)
    f.Println(string(resp))
  }
}
