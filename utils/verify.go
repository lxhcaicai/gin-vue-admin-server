package utils

var (
	LoginVerify    = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	RegisterVerify = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
)
