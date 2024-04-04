package v1

import "github.com/lxhcaicai/gin-vue-admin/server/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
