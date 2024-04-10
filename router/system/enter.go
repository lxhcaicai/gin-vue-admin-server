package system

type RouterGroup struct {
	BaseRouter
	InitRouter
	SysRouter
	UserRouter
	JwtRouter
	OperationRecordRouter
	AuthorityRouter
}
