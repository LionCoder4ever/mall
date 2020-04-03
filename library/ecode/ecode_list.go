package ecode

const IntOk = 0
const IntServerErr = 500
const IntAccountNotFound = 1001

var (
	OK              = add(IntOk)
	ServerErr       = add(IntServerErr)
	AccountNotFound = add(IntAccountNotFound)
)

// internal code description
func init() {
	var defaultMap = make(map[int]string)
	defaultMap[IntOk] = ""
	defaultMap[IntServerErr] = "服务器内部错误"
	defaultMap[IntAccountNotFound] = "未查询到账户"

	Register(defaultMap)
}
