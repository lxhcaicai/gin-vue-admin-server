package system

import "github.com/lxhcaicai/gin-vue-admin/server/global"

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
