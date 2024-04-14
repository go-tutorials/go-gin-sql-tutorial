package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/core-go/config"
	"github.com/core-go/core"
	"github.com/core-go/log"
	"github.com/core-go/log/convert"
	gm "github.com/core-go/log/gin"
	"github.com/core-go/log/strings"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	conf.MiddleWare.Constants = convert.ToCamelCase(conf.MiddleWare.Constants)
	conf.MiddleWare.Map = convert.ToCamelCase(conf.MiddleWare.Map)
	g := gin.New()

	log.Initialize(conf.Log)

	formatter := gm.NewMaskLogger(Mask, Mask)
	logger := gm.NewGinLogger(conf.MiddleWare, log.InfoFields, formatter, MaskLog)

	g.Use(logger.BuildContextWithMask())
	g.Use(logger.Logger())
	g.Use(gin.Recovery())

	err = app.Route(context.Background(), g, conf)
	if err != nil {
		panic(err)
	}

	fmt.Println(core.ServerInfo(conf.Server))
	if err = http.ListenAndServe(core.Addr(conf.Server.Port), g); err != nil {
		fmt.Println(err.Error())
	}
}

func MaskLog(name, s string) string {
	if name == "mobileNo" {
		return strings.Mask(s, 2, 2, "x")
	} else {
		return strings.Mask(s, 0, 5, "x")
	}
}
func Mask(name string, v interface{}) interface{}  {
	if name == "phone" {
		s, ok := v.(string)
		if ok {
			return strings.Mask(s, 0, 3, "*")
		}
	}
	return v
}
