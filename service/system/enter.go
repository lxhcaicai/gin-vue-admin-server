package system

type ServiceGroup struct {
	UserService
	JwtService
	InitDBService
	CasbinService
	OperationRecordService
	SystemConfigService
	AuthorityService
	ApiService
}
