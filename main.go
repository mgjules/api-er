package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/JulesMike/api-er/app"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/security"
	"github.com/JulesMike/api-er/user"
	"github.com/casbin/casbin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/JulesMike/api-er/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AppName is just the app name
const AppName = "API/er"

func main() {
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DB.URI))
	if err != nil {
		logger.Fatal("Can't connect to DB", zap.Error(err))
	}

	db := client.Database(cfg.DB.Name)

	// Services
	userSvc := user.NewService(cfg.Security.PasswordSalt)

	// Repositories
	userRepo := user.NewRepository(db, userSvc)

	// Controllers
	appCtrl := app.NewController()
	securityCtrl := security.NewController(userRepo, userSvc)
	userCtrl := user.NewController(userRepo)

	// Cookie store
	store := cookie.NewStore([]byte(cfg.Security.StoreSecret))

	// Casbin (Authorization)
	enforcer := casbin.NewEnforcer(cfg.Security.Casbin.Model, cfg.Security.Casbin.Policy)

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
	r.Use(security.CSRF(cfg.Security.CSRFSecret, securityCtrl))
	r.Use(user.Auth(enforcer, userSvc, userRepo))

	// Serve public directory
	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	// Gin Routes: Not Found
	r.NoRoute(func(ctx *gin.Context) {
		helper.ResponseNotFound(ctx, "app:route:unknown")
	})

	// Gin Routes
	api := r.Group("/api")
	for _, ctrl := range []Controller{appCtrl, securityCtrl, userCtrl} {
		ctrl.AttachRoutes(api)
	}

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
