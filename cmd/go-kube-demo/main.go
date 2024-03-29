package main

import (
	"context"
	"errors"
	"go-kube-demo/internal/handler"
	"go-kube-demo/internal/pkg/httpclient"
	"go-kube-demo/internal/pkg/logger"
	"go-kube-demo/internal/route"
	"go-kube-demo/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Logger
	logger, err := logger.NewLogger("access.log")
	if err != nil {
		os.Exit(1)
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Next()
		logger.Print(gin.LogFormatterParams{
			Request:  ctx.Request,
			ClientIP: ctx.ClientIP(),
			Method:   ctx.Request.Method,
		})
	})

	// Http client
	httpClient := httpclient.New()

	// Services
	userService := service.NewUserService(httpClient)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()

	// Routes
	routes := route.NewRoutes(router, userHandler, healthHandler)
	routes.AddRoutes()

	// Custom http server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
