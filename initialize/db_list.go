package initialize

import (
	"github.com/lxhcaicai/gin-vue-admin/server/config"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GVA_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}

	if sysDB, ok := dbMap[sys]; ok {
		global.GVA_DB = sysDB
	}
	global.GVA_DBList = dbMap
}
