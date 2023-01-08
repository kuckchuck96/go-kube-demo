package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Health interface {
		Liveness(*gin.Context)
		Readiness(*gin.Context)
	}

	health struct{}

	HealthResponse struct {
		Status string `json:"status"`
		Probe  string `json:"probe"`
	}
)

func NewHealthHandler() Health {
	return &health{}
}

func (health *health) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthResponse{
		Status: "UP",
		Probe:  "liveness",
	})
}

func (health *health) Readiness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthResponse{
		Status: "UP",
		Probe:  "readiness",
	})
}
