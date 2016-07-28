package models

import (
	"github.com/gernest/wuxia/metric"
	"github.com/uber-go/zap"
)

const (
	HCtx = "crossRoutineCtx"
	WCtx = "inAppCtx"
	WCfg = "appConfig"
)

type Context struct {
	Log    zap.Logger
	Metric metric.Metric
	Cfg    *Config
}

type Config struct {
	Port       int
	WorkDir    string
	PublishDir string
}
