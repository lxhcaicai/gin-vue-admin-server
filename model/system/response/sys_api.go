package response

import "github.com/lxhcaicai/gin-vue-admin/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}
