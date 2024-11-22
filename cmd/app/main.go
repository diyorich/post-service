package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"post-storage-service/internal/config"
	"post-storage-service/internal/handler"
	"post-storage-service/internal/repository/pg"
	"post-storage-service/internal/service"
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

	svcManager := service.NewManager(db, cfg.PostProvider)

	r := initRoutes(svcManager)

	fmt.Println("starting server on port ", cfg.App.Port)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	serv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", cfg.App.Port),
		Handler: r,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Println(errors.Wrap(err, "error on starting server"))
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
	h := handler.NewHandler(svcManager.PostService)

	r.GET("/api/posts", h.GetList)
	r.GET("/api/posts/:id", h.GetByID)

	return r
}
