package main

import (
	"net/http"
	"os"

	"github.com/SigiReuvan/iam-service/internal/middleware"
	"github.com/SigiReuvan/iam-service/internal/repository/cache"
	"github.com/SigiReuvan/iam-service/internal/repository/relational"
	"github.com/SigiReuvan/iam-service/internal/service"
	"github.com/SigiReuvan/iam-service/internal/transport"
	"github.com/go-kit/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// TODO: Implement flags or config
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	logger.Log("msg", "starting service")
	defer logger.Log("msg", "stopping service")

	dsn := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log("err", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})

	rep := relational.New(db, logger)
	cache := cache.New(rdb, logger)
	svc := middleware.NewLoggingMiddleware(logger, service.NewService(rep, cache, logger))

	r := transport.NewHttpServer(svc)

	// TODO: Implement Gracefull shutdown
	logger.Log("msg", "starting server", "transport", "http", "addr", "8081")
	logger.Log("err", http.ListenAndServe(":8081", r))
}
