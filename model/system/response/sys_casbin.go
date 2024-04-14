package response

import "github.com/lxhcaicai/gin-vue-admin/server/model/system/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
