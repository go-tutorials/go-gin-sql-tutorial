package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/core-go/config"
	"github.com/core-go/core/server"
	"github.com/core-go/log/convert"
	gm "github.com/core-go/log/gin"
	"github.com/core-go/log/strings"
	"github.com/core-go/log/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/config")
	if err != nil {
		panic(err)
	}
	cfg.MiddleWare.Constants = convert.ToCamelCase(cfg.MiddleWare.Constants)
	cfg.MiddleWare.Map = convert.ToCamelCase(cfg.MiddleWare.Map)
	g := gin.New()

	log.Initialize(cfg.Log)

	formatter := gm.NewMaskLogger(cfg.MiddleWare.Request, Mask, Mask)
	logger := gm.NewGinLogger(cfg.MiddleWare, log.InfoFields, formatter, MaskLog)

	g.Use(logger.BuildContextWithMask())
	g.Use(logger.Logger())
	g.Use(gin.Recovery())

	err = app.Route(context.Background(), g, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(server.ServerInfo(cfg.Server))
	if err = http.ListenAndServe(server.Addr(cfg.Server.Port), g); err != nil {
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
func Mask(obj map[string]interface{}) {
	v, ok := obj["phone"]
	if ok {
		s, ok2 := v.(string)
		if ok2 && len(s) > 3 {
			obj["phone"] = strings.Mask(s, 0, 3, "*")
		}
	}
}
