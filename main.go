package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JulesMike/api-er/controller"
	"github.com/JulesMike/api-er/security"

	"github.com/JulesMike/api-er/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	if cfg.HTTP.CORS.AllowedOrigins != nil && len(cfg.HTTP.CORS.AllowedOrigins) > 0 {
		corsCfg.AllowOrigins = cfg.HTTP.CORS.AllowedOrigins
	} else {
		corsCfg.AllowAllOrigins = true
	}

	// Database
	db, err := gorm.Open(cfg.DB.Dialect, cfg.DB.Name)
	if err != nil {
		logger.Fatal("Can't connect to DB", zap.Error(err))
	}
	defer db.Close()

	autoMigrate(db)

	controller.Init(db)

	security.Init(cfg.Security.PasswordSalt)

	// Cookie store
	store := cookie.NewStore([]byte(cfg.Security.StoreSecret))

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
	r.Use(sessions.Sessions(cfg.Security.SessionKey, store))

	// Static Routes Middlewares
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	attachRoutes(r)

	// Run gin server
	url := cfg.HTTP.Host + ":" + cfg.HTTP.Port
	logger.Info("Gin server started on ", zap.String("url", url))
	s := &http.Server{
		Addr:           url,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
