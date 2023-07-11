package logger

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"resume-resolving/internal/app/service/resume/config"
	"resume-resolving/internal/pkg/logger"
)

const (
	modeDev = "dev"
)

type KitexLogrus struct {
	config *config.Config
}

func (k *KitexLogrus) Init() error {
	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   k.config.ConfigInNacos.Log.FilePath,
		MaxSize:    k.config.ConfigInNacos.Log.MaxSize,
		MaxBackups: k.config.ConfigInNacos.Log.MaxBackups,
		MaxAge:     k.config.ConfigInNacos.MaxAge,
		Compress:   true,
	}
	klog.SetLogger(logrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	if k.config.ConfigInNacos.Mode == modeDev {
		klog.SetOutput(os.Stdout)
		klog.SetOutput(lumberjackLogger)
	} else {
		klog.SetOutput(lumberjackLogger)
	}
	return nil
}

func NewKitexLogrus(config *config.Config) logger.Logger {
	return &KitexLogrus{
		config: config,
	}
}
