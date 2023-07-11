package logger

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"resume-resolving/internal/app/web/config"
	"resume-resolving/internal/pkg/logger"
)

type HertzLogrus struct {
	config *config.Config
}

func (h *HertzLogrus) Init() error {
	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   h.config.ConfigInNacos.Log.FilePath,
		MaxSize:    h.config.ConfigInNacos.Log.MaxSize,    // 一个文件最大可达20M。
		MaxBackups: h.config.ConfigInNacos.Log.MaxBackups, // 最多同时保存 5 个文件。
		MaxAge:     h.config.ConfigInNacos.Log.MaxAge,     // 一个文件最多可以保存 10 天。
		Compress:   true,                                  // 用 gzip 压缩。
	}
	logge := logrus.NewLogger()
	logge.SetOutput(lumberjackLogger)
	logge.SetLevel(hlog.LevelDebug)
	hlog.SetLogger(logge)
	return nil
}

func NewHertzLogrus(config *config.Config) logger.Logger {
	return &HertzLogrus{
		config: config,
	}
}
