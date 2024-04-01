package internal

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"go.uber.org/zap/zapcore"
	"os"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct {
}

func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.GVA_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
