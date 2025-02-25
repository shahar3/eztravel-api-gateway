package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shahar3/api-gateway/config"
	middlewares "github.com/shahar3/api-gateway/middleware"
	"github.com/shahar3/api-gateway/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Initialize Gin engine
	router := gin.New()

	// Global middleware
	router.Use(middlewares.Logger(logger))
	router.Use(middlewares.Recovery(logger))

	// Set timeouts and other HTTP server settings
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	logger.Infof("Starting API Gateway on port %s in %s mode", cfg.Port, cfg.Env)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server failed: %v", err)
	}
}
