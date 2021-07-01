package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rpsingh21/checklist-api/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	config := config.NewConfig()

	app := ConfigApp(sugar, config)
	go app.Run()

	sdc := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(sdc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	sig := <-sdc
	sugar.Infof("Got interupt signal %v", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.Shutdown(&ctx)
}
