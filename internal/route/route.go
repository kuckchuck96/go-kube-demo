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
		engine        *gin.Engine
		userHandler   handler.User
		healthHandler handler.Health
	}
)

func NewRoutes(
	engine *gin.Engine,
	userHandler handler.User,
	healthHandler handler.Health,
) Route {
	return &route{
		engine,
		userHandler,
		healthHandler,
	}
}

func (route *route) AddRoutes() {
	group := route.engine.Group("/go-kube-demo")
	{
		addUserRoutes(route.userHandler, group)
		addHealhRoutes(route.healthHandler, group)
	}
}

func addUserRoutes(userHandler handler.User, group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	userGroup.GET("/all", userHandler.GetUsers)
}

func addHealhRoutes(healthHandler handler.Health, group *gin.RouterGroup) {
	healthGroup := group.Group("/health")
	healthGroup.GET("/liveness", healthHandler.Liveness)
	healthGroup.GET("/readiness", healthHandler.Readiness)
}
