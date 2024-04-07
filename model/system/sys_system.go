package system

import "github.com/lxhcaicai/gin-vue-admin/server/config"

type System struct {
	Config config.Server `json:"config"`
}
