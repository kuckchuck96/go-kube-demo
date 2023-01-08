package route

import (
	"go-kube-demo/internal/handler"

	"github.com/gin-gonic/gin"
)

type (
	Route interface {
		AddRoutes()
	}

	route struct {
		engine      *gin.Engine
		userHandler handler.User
	}
)

func NewRoutes(
	engine *gin.Engine,
	userHandler handler.User,
) Route {
	return &route{
		engine,
		userHandler,
	}
}

func (route *route) AddRoutes() {
	group := route.engine.Group("/go-kube-demo")
	{
		addUserRoutes(route, group)
	}
}

func addUserRoutes(route *route, group *gin.RouterGroup) {
	userGroup := group.Group("user")
	userGroup.GET("all", route.userHandler.GetUsers)
}
