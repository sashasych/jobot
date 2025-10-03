package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"jobot/internal/application"
	"jobot/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	log := logger.InitProdLogger()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	appl, err := application.NewApplication(log)
	if err != nil {
		cancel()
		log.Fatal("Failed to create application", zap.Error(err))
	}

	wg := sync.WaitGroup{}

	appl.Start(ctx, &wg, cancel)

	wg.Wait()

	cancel()
}
