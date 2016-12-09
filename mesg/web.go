package mesg

import "github.com/jellybean4/gosalt/models"

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
		models.Template
	}

	SetTemplateRep struct {
		Result
	}

	GetTemplateReq struct {
		Name string `json:"name"`
	}

	GetTemplateRep struct {
		Result
		models.Template
	}

	DelTemplateReq struct {
		Name string `json:"name"`
	}

	DelTemplateRep struct {
		Result
	}
)
