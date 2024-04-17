package example

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	FileUploadAndDownloadApi
	CustomerApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
)
