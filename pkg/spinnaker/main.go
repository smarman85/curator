package spinnaker

import (
  f "fmt"
  "encoding/json"
)

type pipeline struct {
  Schema string `json:"schema"`
  Name string `json:"name"`
  Application string `json:"application"`
  Template tempInfo `json:"template"`
  Vars variables `json:"variables"`
}

type tempInfo struct {
  ArtifactAccount string `json:"artifactAccount"`
  Reference string `json:"reference"`
  TempType string `json:"type"`
}

type variables struct {
  Chart string `json:"chartVersion"`
  Repo string `json:"repo_name"`
}

func genPipeline(app, repo string) []byte {
  pipe := pipeline{
    Schema: "v2",
    Name: "test-" + app,
    Application: app,
    Template: tempInfo{
      ArtifactAccount: "front50ArtifactCredentials",
      Reference: "spinnaker://onerails:latest",
      TempType: "front50/pipelineTemplate",
    },
    Vars: variables{
      Chart: "THISisAChartVersion",
      Repo: repo,
    },
  }

  p, err := json.MarshalIndent(&pipe, "", "    ")
  if err != nil {
    f.Println("Error marshalling pipeline", err)
  }

  return p
}

func Pipeline (app, repo string) {
  f.Println("Inside spinnaker package")
  f.Println(app)
  pipeline := genPipeline(app, repo)
  f.Println(string(pipeline))
}
