package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	serverAddress := ":8000"
	args := os.Args
	if len(args) > 1 {
		serverAddress = args[1]
	}
	defer logger.Sync()

	router := mux.NewRouter()

	server := http.Server{
		Addr:    serverAddress, // configure the bind address
		Handler: router,        // set the default handler
		// ErrorLog:     (*log.Logger)(logger.Sugar()), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	sugar := logger.Sugar()
	go func() {
		sugar.Infof("Starting server on %s", serverAddress)
		if err := server.ListenAndServe(); err != nil {
			sugar.Debugf("Shutdown server due to %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	sugar.Infof("Got interupt signal %v", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)

}
