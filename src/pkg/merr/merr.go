package merr

var Msgs = map[int] string {
	ERROR:    		"fail",
	SUCCESS: 		"success",
	LOGIN_ERROR: 	"login failed",
}

func GetMsg(code int) string {
	msg, ok := Msgs[code]
	if ok {
		return msg
	}

	return Msgs[ERROR]
}