package system

type ServiceGroup struct {
	UserService
	JwtService
	InitDBService
	CasbinService
	OperationRecordService
}
