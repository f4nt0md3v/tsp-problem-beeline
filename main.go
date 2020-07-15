package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/f4nt0md3v/tsp-problem-beeline/proto/route"
	"github.com/f4nt0md3v/tsp-problem-beeline/server"
)

func main() {
	// Initialize logger
	logger, _ := zap.NewDevelopment()
	defer func() {
		if err := logger.Sync(); err != nil { // flushes buffer, if any
			log.Println(err)
		}
	}()

	// Initialize and register everything
	sugar := logger.Sugar()
	gSrv := grpc.NewServer()
	srv := server.NewRouteServer(sugar)
	route.RegisterRouteServer(gSrv, srv)
	reflection.Register(gSrv)

	listener, err := net.Listen("tcp", ":9090") // nolint:gosec
	if err != nil {
		sugar.Fatalf("Unable to listen: %v", err)
	}

	err = gSrv.Serve(listener)
	if err != nil {
		sugar.Fatalf("Unable to start gRPC server: %v", err)
	}

	sugar.Info("gRPC service has started")
}
