package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func Route(g *gin.Engine, ctx context.Context, root Root) error {
	app, err := NewApp(ctx, root)
	if err != nil {
		return err
	}

	// g.GET("/health", app.HealthHandler.Check)

	 userPath := g.Group("/users")
	 {
		 userPath.GET("", app.UserHandler.GetAll)
		 userPath.GET("/:id", app.UserHandler.Load)
		 userPath.POST("", app.UserHandler.Insert)
		 userPath.PUT("/:id", app.UserHandler.Update)
		 userPath.PATCH("/:id", app.UserHandler.Patch)
		 userPath.DELETE("/:id", app.UserHandler.Delete)
	}

	//r.HandleFunc(userPath+"/{id}", app.UserHandler.Load).Methods(GET)
	//r.HandleFunc(userPath, app.UserHandler.Insert).Methods(POST)
	//r.HandleFunc(userPath+"/{id}", app.UserHandler.Update).Methods(PUT)
	//r.HandleFunc(userPath+"/{id}", app.UserHandler.Patch).Methods(PATCH)
	//r.HandleFunc(userPath+"/{id}", app.UserHandler.Delete).Methods(DELETE)

	return nil
}
