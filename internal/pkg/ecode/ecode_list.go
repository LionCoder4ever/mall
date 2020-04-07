package ecode

const (
	intOk           = 0
	intParameterErr = 1
	intServerErr    = 500

	// account service
	intAccountNotFound        = 1001
	intPhoneHasBeenRegistered = 1002
	intPwdRepeatCheckErr      = 1003
	intSavePwdErr             = 1004
)

var (
	OK                     = add(intOk)
	ParameterErr           = add(intParameterErr)
	ServerErr              = add(intServerErr)
	AccountNotFound        = add(intAccountNotFound)
	PhoneHasBeenRegistered = add(intPhoneHasBeenRegistered)
	PwdRepeatCheckErr      = add(intPwdRepeatCheckErr)
	SavePwdErr             = add(intSavePwdErr)
)

// internal code description
func init() {
	var defaultMap = make(map[int]string)
	defaultMap[intOk] = ""
	defaultMap[intParameterErr] = "check http request parameter"
	defaultMap[intServerErr] = "服务器内部错误"
	defaultMap[intAccountNotFound] = "未查询到账户"
	defaultMap[intPhoneHasBeenRegistered] = "phone has been registered"
	defaultMap[intPwdRepeatCheckErr] = "password check fail"
	defaultMap[intSavePwdErr] = "please register later"

	Register(defaultMap)
}
