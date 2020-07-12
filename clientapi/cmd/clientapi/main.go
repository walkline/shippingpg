package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/clientapi/kvimporting/json"
	"github.com/walkline/shippingpg/clientapi/port"
	"github.com/walkline/shippingpg/clientapi/port/pb"
	"github.com/walkline/shippingpg/clientapi/server"
	"google.golang.org/grpc"
)

func main() {
	// TODO: rewrite to config loader
	port := "8080"
	if p := os.Getenv("HTTP_PORT"); len(p) > 0 {
		port = p
	}

	portGRPCAddr := "localhsot:6969"
	if p := os.Getenv("PORT_GRPC_ADDR"); len(p) > 0 {
		portGRPCAddr = p
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	conn, portClient := portGRPCClient(portGRPCAddr, logger)
	defer conn.Close()

	// handling graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		logger.Log("received syscall", oscall)
		cancel()
	}()

	if err := serve(ctx, port, logger, portClient); err != nil {
		logger.Log("failed to serve:", err)
	}
}

func serve(ctx context.Context, p string, logger log.Logger, portClient pb.PortDomainServiceClient) error {
	portsRepo := port.NewKVRepository(portClient)
	srv := server.NewHTTPServer(
		":"+p,
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

func portGRPCClient(addr string, logger log.Logger) (*grpc.ClientConn, pb.PortDomainServiceClient) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		err := fmt.Sprintf("fail to dial: %s", err.Error())
		logger.Log("grpc_dial", err)
		panic(err)
	}
	return conn, pb.NewPortDomainServiceClient(conn)
}
