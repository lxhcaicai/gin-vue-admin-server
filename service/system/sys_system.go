package system

import (
	"github.com/lxhcaicai/gin-vue-admin/server/config"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	"github.com/lxhcaicai/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOs()
	if s.Cpu, err = utils.InitCpu(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

func (systemConfigService *SystemConfigService) GetSystemConfig() (config config.Server, err error) {
	return global.GVA_CONFIG, nil
}

// SetSystemConfig
//
//	@Description: 设置配置文件
func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
