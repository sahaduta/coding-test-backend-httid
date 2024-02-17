package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sahaduta/coding-test-backend-httid/database"
	"github.com/sahaduta/coding-test-backend-httid/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := database.NewConn()
	if err != nil {
		log.Fatal("fail to connect to database")
	}

	r := server.NewRouter(server.GetRouterOpts(db))

	srv := http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	secondValue, _ := strconv.Atoi(os.Getenv("SHUTDOWN_TIME"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(secondValue)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
