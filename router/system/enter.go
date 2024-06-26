package system

type RouterGroup struct {
	BaseRouter
	InitRouter
	SysRouter
	UserRouter
	JwtRouter
	OperationRecordRouter
	AuthorityRouter
	CasbinRouter
	ApiRouter
	DictionaryDetailRouter
	AuthorityBtnRouter
	MenuRouter
	SysExportTemplateRouter
}
