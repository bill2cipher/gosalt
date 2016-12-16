package mesg

type SaltLoginRep struct {
	Pems   []string `json:"perms"`
	Start  float32  `json:"start"`
	Expire float32  `json:"expire"`
	Token  string   `json:"token"`
	User   string   `json:"user"`
	EAuth  string   `json:"eauth"`
}

type SaltLoginReq struct {
	User  string `json:"username"`
	Pass  string `json:"password"`
	EAuth string `json:"eauth"`
}

type SaltExecRep struct {
}
