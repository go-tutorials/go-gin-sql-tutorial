package app

import (
	"github.com/core-go/core/server"
	"github.com/core-go/log/gin"
	"github.com/core-go/log/zap"
	"github.com/core-go/sql"
)

type Config struct {
	Server     server.ServerConf `mapstructure:"server"`
	Sql        sql.Config        `mapstructure:"sql"`
	Log        log.Config        `mapstructure:"log"`
	MiddleWare gin.LogConfig     `mapstructure:"middleware"`
}
