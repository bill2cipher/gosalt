package mesg

const (
	SUCCESS_MESG = ""
	SUCCESS_CODE = 0
)

const (
	GIN_PARSE_REQUEST_LOG  = "gin read request failed"
	GIN_PARSE_REQUEST_CODE = 1
)

const (
	CACHE_GET_DATA_LOG  = "cache get failed"
	CACHE_GET_DATA_CODE = 2
)

const (
	CACHE_GET_TYPE_LOG  = "type assert failed"
	CACHE_GET_TYPE_CODE = 3
)

const (
	CACHE_SET_DATA_LOG  = "cache set failed"
	CACHE_SET_DATA_CODE = 4
)

const (
	CACHE_DEL_DATA_LOG  = "cache del failed"
	CACHE_DEL_DATA_CODE = 5
)

const (
	JSON_UNMARSHAL_LOG  = "json unmarshal failed"
	JSON_UNMARSHAL_CODE = 6
)

const (
	TYPE_ASSERT_LOG  = "type assert failed"
	TYPE_ASSERT_CODE = 7
)
