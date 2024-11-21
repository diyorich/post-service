package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"posts-service/internal/config"
	"posts-service/internal/repository/pg"
	"posts-service/internal/service"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	//connect to db
	db, err := pg.Dial(*cfg)
	if err != nil {
		return err
	}

	svcManager := service.NewManager(db)

	r := initRoutes(svcManager)

	fmt.Printf("starting server on port %v", cfg.App.Port)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	serv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", cfg.App.Port),
		Handler: r,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			fmt.Println("error on starting server")
		}
	}()

	fmt.Println("server started")
	<-done
	fmt.Println("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		fmt.Println("failed on stopping server")
		return nil
	}

	fmt.Println("server stopped")
	return nil
}

func initRoutes(svcManager *service.Manager) *gin.Engine {
	r := gin.Default()

	return r
}
