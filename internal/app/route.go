package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Route(g *gin.Engine, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}

	userPath := g.Group("/users")
	{
		userPath.GET("", app.UserHandler.All)
		userPath.GET("/:id", app.UserHandler.Load)
		userPath.POST("", app.UserHandler.Insert)
		userPath.PUT("/:id", app.UserHandler.Update)
		userPath.PATCH("/:id", app.UserHandler.Patch)
		userPath.DELETE("/:id", app.UserHandler.Delete)
	}

	return nil
}
