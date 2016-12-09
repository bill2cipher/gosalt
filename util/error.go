package util

const (
	SUCCESS_MESG = ""
	SUCCESS_CODE = 0
)

const (
	GIN_PARSE_REQUEST_LOG  = "gin read request"
	GIN_PARSE_REQUEST_CODE = 1
)

const (
	CACHE_GET_DATA_LOG  = "cache get"
	CACHE_GET_DATA_CODE = 2

	CACHE_GET_TYPE_LOG  = "type assert"
	CACHE_GET_TYPE_CODE = 3

	CACHE_SET_DATA_LOG  = "cache set"
	CACHE_SET_DATA_CODE = 4

	CACHE_DEL_DATA_LOG  = "cache del"
	CACHE_DEL_DATA_CODE = 5
  
  CACHE_INIT_DATA_LOG = "cache init"
  CACHE_INIT_DATA_CODE = 16
)

const (
	JSON_MARSHAL_LOG    = "json marshal"
	JSON_MARSHAL_CODE   = 15
	JSON_UNMARSHAL_LOG  = "json unmarshal"
	JSON_UNMARSHAL_CODE = 6
)

const (
	TYPE_ASSERT_LOG  = "type assert"
	TYPE_ASSERT_CODE = 7
)

const (
	TEMPL_PARSE_LOG  = "template parse"
	TEMPL_PARSE_CODE = 8
)

const (
	TEMPL_EXEC_LOG  = "template execute"
	TEMPL_EXEC_CODE = 9
)

const (
	FILE_OPEN_LOG  = "file open"
	FILE_OPEN_CODE = 10
)
const (
	FILE_CLOSE_LOG  = "file close"
	FILE_CLOSE_CODE = 11
)

const (
	FILE_DELETE_LOG  = "file delete"
	FILE_DELETE_CODE = 12
)

const (
	SCRIPT_EXECUTE_LOG  = "script execute"
	SCRIPT_EXECUTE_CODE = 13
)

const (
	DIR_CREAE_LOG  = "directory create"
	DIR_CREAE_CODE = 14
)

const (
	DB_CONN_LOG   = "db connection"
	DB_CREATE_LOG = "db create table"
	DB_FETCH_LOG  = "db fetch"
	DB_STORE_LOG  = "db store"
	DB_DELETE_LOG = "db delete"
	DB_ALL_LOG    = "db fetch all"
)
