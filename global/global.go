package global

import (
	"github.com/lxhcaicai/gin-vue-admin/server/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server

	BlackCache local_cache.Cache
)
