package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/datastore"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/handler"
	"github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web/service"
)

var (
	ShutdownTimeoutSecond = 5 * time.Second
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Extract configurations from the environment variables
	config := handler.ExtractConfig()

	log.Info("starting server...")

	// Initialize a in-memory bookService instance
	sv := service.NewInMemBookService(datastore.NewInMemStore())

	// Create a HTTP router object which contains a gin.Engine
	// The default gin.Engine has the Logger and Recovery middleware attached
	router := handler.NewHandler(gin.Default(), config, sv)
	router.CreateEndpoints()

	srv := &http.Server{
		Addr:    config.Port,
		Handler: router.Engine,
	}

	log.WithField("Port", config.Port).Info("Listening on port")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeoutSecond)
	defer cancel()

	// Shutdown server
	log.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
