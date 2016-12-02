package mesg

type Result struct {
  Code int    `json:"code"`
  Mesg string `json:"mesg"`
}

type ReleaseReq struct {
  Version string   `json:"version"`
  Types   []string `json:"types"`
}

type ReleaseRep struct {
  Result
  Pacakge string `json:"package"`
}

type DeployReq struct {
  Server  string `json:"string"`
  Version string `json:"string"`
}

type DeployRep struct {
  Result
}
