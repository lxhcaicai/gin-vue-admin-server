package response

import "github.com/lxhcaicai/gin-vue-admin/server/model/example"

type FileResponse struct {
	File example.ExaFile `json:"file"`
}

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}
