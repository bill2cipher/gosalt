package db

var (
	SERVER_TABLE   = []byte("server")
	TEMPLATE_TABLE = []byte("template")
	CONFIGS_TABLE  = []byte("configs")

	TABLES = [][]byte{SERVER_TABLE, TEMPLATE_TABLE, CONFIGS_TABLE}
)
