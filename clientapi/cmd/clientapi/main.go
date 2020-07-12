package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/clientapi/inmem"
	"github.com/walkline/shippingpg/clientapi/kvimporting/json"
	"github.com/walkline/shippingpg/clientapi/server"
)

func main() {
	// TODO: rewrite to config loader
	port := "8080"
	if p := os.Getenv("HTTP_PORT"); len(p) > 0 {
		port = p
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// handling graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		logger.Log("received syscall", oscall)
		cancel()
	}()

	if err := serve(ctx, port, logger); err != nil {
		logger.Log("failed to serve:", err)
	}
}

func serve(ctx context.Context, port string, logger log.Logger) error {
	portsRepo := inmem.NewKeyValueRepo()
	srv := server.NewHTTPServer(
		":"+port,
		log.With(logger, "module", "http"),
		json.NewReaderImporter(portsRepo),
		portsRepo,
	)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Log("listening failed", err)
		}
	}()

	logger.Log("HTTP", "Server is running...")

	<-ctx.Done()

	logger.Log("HTTP", "Shutting down HTTP Server...")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctxShutDown)
	if err != nil {
		logger.Log("failed", err)
	}

	logger.Log("HTTP", "Server stopped")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}
