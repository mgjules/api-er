package main

import (
	"log"
	"time"

	"github.com/JulesMike/api-er/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AppName is just the app name
const AppName = "API/er"

func main() {
	var err error

	// Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("[%s] Can't load config : %v", AppName, err)
	}

	// Logger
	var logger *zap.Logger
	if cfg.Prod {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("[%s] Can't create logger : %v", AppName, err)
	}
	defer logger.Sync()

	// Cors
	corsCfg := cors.DefaultConfig()
	if cfg.CORS.AllowedOrigins != nil && len(cfg.CORS.AllowedOrigins) > 0 {
		corsCfg.AllowOrigins = cfg.CORS.AllowedOrigins
	} else {
		corsCfg.AllowAllOrigins = true
	}

	// Gin server
	if cfg.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global middlewares
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(cors.New(corsCfg))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	attachRoutes(r)

	// Run gin server
	url := cfg.HTTP.Host + ":" + cfg.HTTP.Port
	logger.Info("Gin server started on ", zap.String("url", url))
	r.Run(url)
}
