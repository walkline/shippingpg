package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/portdomain/inmem"
	"github.com/walkline/shippingpg/portdomain/server"
	"github.com/walkline/shippingpg/portdomain/server/pb"
	"google.golang.org/grpc"
)

func main() {
	// TODO: rewrite to config loader
	port := "6969"
	if p := os.Getenv("GRPC_PORT"); len(p) > 0 {
		port = p
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		logger.Log("received syscall", oscall)
		cancel()
	}()

	if err := serve(ctx, port, logger); err != nil {
		logger.Log("serv_err", err)
	}
}

func serve(ctx context.Context, port string, logger log.Logger) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Log("listening_err", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPortDomainServiceServer(
		grpcServer,
		server.NewPortServiceServer(inmem.NewPortRepository(), logger),
	)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
		}
	}()

	logger.Log("GRPC", "Server is running...")

	<-ctx.Done()

	logger.Log("GRPC", "Shutting down GRPC Server...")

	grpcServer.GracefulStop()

	logger.Log("GRPC", "Server stopped")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}
