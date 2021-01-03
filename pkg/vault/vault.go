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

func VaultApi(url, app string) int {
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

func createVaultEpt(url string, data []byte) []byte {
  request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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

func EndpointExists(app string) bool {
  rcode := VaultApi(vltUrl, app)
  exists := false
  if rcode == 200 {
    exists = true
  }
  return exists
}

func Endpoint(app, repo string) {
  if !EndpointExists(app) {
    jData, _ := json.Marshal(map[string]map[string]string{
      "data": {"env":"test"},
    })
    f.Println("Creating Endpoint for: " + app)
    resp := createVaultEpt(vltUrl + "v1/secret/data/test/" + app + "/env", jData)
    f.Println(string(resp))
    /*
    bot, _ := json.Marshal(map[string]map[string]string{
      "data": {app: repo},
    })
    createVaultEpt(vltUrl + "v1/secret/data/test/apps", bot)
    */
  }
}
