package global

import (
	"github.com/lxhcaicai/gin-vue-admin/server/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	BlackCache local_cache.Cache
)
