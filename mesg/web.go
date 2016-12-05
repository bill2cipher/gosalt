package mesg

type Result struct {
	Code int    `json:"code"`
	Mesg string `json:"mesg"`
}

type (
	ReleaseReq struct {
		Version string   `json:"version"`
		Types   []string `json:"types"`
	}

	ReleaseRep struct {
		Result
		Pacakge string `json:"package"`
	}
)

type (
	DeployReq struct {
		Server  string `json:"string"`
		Version string `json:"string"`
	}

	DeployRep struct {
		Result
	}
)

type (
	SetTemplateReq struct {
		Name   string            `json:"name"`
		Config map[string]string `json:"config"`
	}

	SetTemplateRep struct {
		Result
	}

	GetTemplateReq struct {
		Name string `json:"name"`
	}

	GetTemplateRep struct {
		Result
		Name   string            `json:"name"`
		Config map[string]string `json:"config"`
	}

	DelTemplateReq struct {
		Name string `json:"name"`
	}

	DelTemplateRep struct {
		Result
	}
)
