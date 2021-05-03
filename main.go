package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cerncorp/larago/app/http/middlewares"
	"github.com/cerncorp/larago/config"
	"github.com/cerncorp/larago/routes"
	"github.com/gin-gonic/gin"
)

func setupLogOutput() (*os.File, error) {
	f, err := os.Create("./storage/logs/gin.log")

	if err == nil {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	return f, err
}

func main() {
	fmt.Println("Larago Hello!")

	// open gin.log
	f, err := setupLogOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// create *gin.engine
	r := gin.New()

	// add middlewares: recovery panic-> 1 route, logger, basic auth
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth()) //, gin.Logger()

	// add routes
	r, _ = routes.LoadRoutes(r)

	_address := config.Cfg.Server.Host + ":" + config.Cfg.Server.Port
	srv := &http.Server{
		Addr:    _address,
		Handler: r,
	}

	// run server
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			return
		}
	}()

	// grateful shutdown
	// Wait for interupt signal to fracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown: ", err)
	}

	fmt.Println("Server exiting")
}
