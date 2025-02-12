package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SigiReuvan/iam-service/cmd/internal/util"
	"github.com/SigiReuvan/iam-service/config"
	"github.com/SigiReuvan/iam-service/internal/middleware"
	"github.com/SigiReuvan/iam-service/internal/repository/cache"
	"github.com/SigiReuvan/iam-service/internal/repository/relational"
	"github.com/SigiReuvan/iam-service/internal/service"
	"github.com/SigiReuvan/iam-service/internal/transport"
	"github.com/go-kit/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func main() {
	// Initialize logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	logger.Log("msg", "starting service")

	// Load configuration
	cfg := config.Load(logger)
	gormCfg := &gorm.Config{}
	gormCfg.Logger = gormLogger.Default.LogMode(gormLogger.Silent)

	// Connect to PostgreSQL
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
	db, err := gorm.Open(postgres.Open(dsn), gormCfg)

	stopSvcMsg := "stopping service"

	if err != nil {
		errStr := util.FormatErrorToString(err)
		logger.Log("err", "failed to connect to postgres", "detail", errStr)
		logger.Log("msg", stopSvcMsg)
		os.Exit(1)
	}

	// Get the underlying sql.DB to perform further operations and graceful shutdown.
	sqlDB, err := db.DB()
	if err != nil {
		errStr := util.FormatErrorToString(err)
		logger.Log("err", "failed to retrieve sql.DB from gorm", "detail", errStr)
		logger.Log("msg", stopSvcMsg)
		os.Exit(1)
	}
	// Ping PostgreSQL to ensure the connection is healthy.
	if err = sqlDB.Ping(); err != nil {
		errStr := util.FormatErrorToString(err)
		logger.Log("err", "failed to ping postgres", "detail", errStr)
		logger.Log("msg", stopSvcMsg)
		os.Exit(1)
	}

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})
	// Ping Redis to ensure the connection is healthy.
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		errStr := util.FormatErrorToString(err)
		logger.Log("err", "failed to ping redis", "detail", errStr)
		logger.Log("msg", stopSvcMsg)
		os.Exit(1)
	}

	// Initialize repositories and services
	rep := relational.New(db, logger)
	cacheRepo := cache.New(rdb, logger)
	svc := middleware.NewLoggingMiddleware(logger, service.NewService(rep, cacheRepo, logger))
	handler := transport.NewHttpServer(svc)

	// Create the HTTP server
	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	// Start the HTTP server in a separate goroutine.
	go func() {
		logger.Log("msg", "starting server", "transport", "http", "addr", "8081")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log("err", "HTTP server error", "detail", err)
		}
	}()

	// Create channel to listen for interrupt or terminate signals.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Log("msg", "shutting down server...")

	// Create a context with timeout for the graceful shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log("err", "server forced to shutdown", "detail", err)
	}

	// Close the Postgres connection gracefully.
	if err := sqlDB.Close(); err != nil {
		logger.Log("err", "failed to close postgres connection", "detail", err)
	}

	// Close the Redis connection gracefully.
	if err := rdb.Close(); err != nil {
		logger.Log("err", "failed to close redis connection", "detail", err)
	}

	logger.Log("msg", "server exiting")
}
