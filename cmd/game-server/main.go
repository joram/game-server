package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joram/game-server/config"
	"github.com/joram/game-server/pb"
	"github.com/joram/game-server/server"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	svc, err := server.New()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	pb.RegisterGameServerServer(grpcServer, svc)
	log.Print("Starting game-server GRPC server")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
