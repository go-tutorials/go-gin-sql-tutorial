package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/core-go/config"
	sv "github.com/core-go/service"
	"github.com/gin-gonic/gin"
	"net/http"

	"go-service/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	g := gin.New()

	g.Use(gin.Logger())

	g.Use(gin.Recovery())

	g.Use(ginBodyLogMiddleware())

	er2 := app.Route(g , context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println(sv.ServerInfo(conf.Server))
	if er3 := http.ListenAndServe(sv.Addr(conf.Server.Port), g); er3 != nil {
		fmt.Println(er3.Error())
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString("\n"), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		fmt.Println("Response body: " + blw.body.String())
	}
}
