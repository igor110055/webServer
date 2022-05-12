package config

var (
	// API root Group
	API = "/api"

	// V1 v1 Group
	V1           = "/v1"
	API_V1_GROUP = API + V1

	// POOL Group
	POOL = "/pool"

	// COMMON common
	COMMON = "/common"

	SUCCESS = int64(2000)
	ERROR   = int64(5000)

	INVALID_PARAMS = int64(4000)
	INTERNAL_ERROR = int64(4001)
	DBERROR        = int64(4002)
	GETERROR       = int64(4003)
)

var MsgFlags = map[int64]string{
	SUCCESS:        "Success",
	ERROR:          "Fail",
	INTERNAL_ERROR: "Internal Error",
	INVALID_PARAMS: "Request Parameter Error",
	DBERROR:        "Update Database Error",
	GETERROR:       "Get Price Error",
}

func GetMsg(code int64) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
