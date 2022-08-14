package app

import (
	"github.com/core-go/core"
	"github.com/core-go/log"
	"github.com/core-go/sql"

	"go-service/pkg/gin"
)

type Config struct {
	Server     core.ServerConf `mapstructure:"server"`
	Sql        sql.Config      `mapstructure:"sql"`
	Log        log.Config      `mapstructure:"log"`
	MiddleWare gin.LogConfig   `mapstructure:"middleware"`
}
