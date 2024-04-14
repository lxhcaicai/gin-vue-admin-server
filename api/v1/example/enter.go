package example

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
