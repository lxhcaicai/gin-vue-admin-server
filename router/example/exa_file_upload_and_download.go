package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
)

type FileUploadAndDownloadRouter struct {
}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		FileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)
		FileUploadAndDownloadRouter.POST("getFileList", exaFileUploadAndDownloadApi.GetFileList)
		FileUploadAndDownloadRouter.POST("deleteFile", exaFileUploadAndDownloadApi.DeleteFile)     // 删除指定文件
		FileUploadAndDownloadRouter.POST("editFileName", exaFileUploadAndDownloadApi.EditFileName) // 编辑文件名或者备注

	}
}
