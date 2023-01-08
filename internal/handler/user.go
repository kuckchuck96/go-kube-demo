package handler

import (
	"go-kube-demo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	User interface {
		GetUsers(*gin.Context)
	}

	user struct {
		userService service.User
	}
)

func NewUserHandler(userService service.User) User {
	return &user{
		userService,
	}
}

func (user *user) GetUsers(ctx *gin.Context) {
	users, err := user.userService.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
