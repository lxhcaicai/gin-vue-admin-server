package response

import "github.com/lxhcaicai/gin-vue-admin/server/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}
