package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sahaduta/coding-test-backend-httid/database"
	"github.com/sahaduta/coding-test-backend-httid/pkg/logger"
	"github.com/sahaduta/coding-test-backend-httid/server"
)

func main() {
	logger.SetLogrusLogger()
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatalf(err.Error())
	}

	db, err := database.NewConn()
	if err != nil {
		logger.Log.Fatal("fail to connect to database")
	}

	r := server.NewRouter(server.GetRouterOpts(db))

	srv := http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Infof("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Infof("Shutting down server...")

	secondValue, _ := strconv.Atoi(os.Getenv("SHUTDOWN_TIME"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(secondValue)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown:", err)
	}

	logger.Log.Infof("Server exiting")
}
