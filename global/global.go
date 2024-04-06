package global

import (
	"github.com/lxhcaicai/gin-vue-admin/server/config"
	"github.com/lxhcaicai/gin-vue-admin/server/utils/timer"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_VP                  *viper.Viper
	GVA_DBList              map[string]*gorm.DB
	GVA_CONFIG              config.Server
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	BlackCache              local_cache.Cache
	GVA_REDIS               *redis.Client
	GVA_Concurrency_Control = &singleflight.Group{}
)
