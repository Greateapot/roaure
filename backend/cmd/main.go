package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Greateapot/roaure/internal/database"
	roaurev1 "github.com/Greateapot/roaure/internal/genproto/roaure/v1"
	"github.com/Greateapot/roaure/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	configPath := os.Getenv("CONFIG_PATH")
	network := os.Getenv("SERVER_NETWORK")
	address := os.Getenv("SERVER_ADDRESS")
	ledChip := os.Getenv("LED_CHIP")
	rawLedlineOffset := os.Getenv("LED_LINE_OFFSET")

	ledLineOffset, err := strconv.ParseInt(rawLedlineOffset, 10, 64)
	if err != nil {
		grpclog.Fatal(err)
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		grpclog.Fatal(err)
	}

	database := database.NewDatabase(configPath)

	s := grpc.NewServer()
	roaurev1.RegisterRoaureServiceServer(s, server.NewRoaureServiceServer(ctx, database, ledChip, int(ledLineOffset)))
	reflection.Register(s)

	// Красивая остановка по сигналу
	go GracefulStop(s)

	if err := s.Serve(lis); err != nil {
		grpclog.Fatal(err)
	}
}

func GracefulStop(s *grpc.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	grpclog.Infoln("shutting down...")
	s.GracefulStop()
	grpclog.Infoln("done")
}
