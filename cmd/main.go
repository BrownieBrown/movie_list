package main

import (
	"movie_list/pkg/protobuf"
	"movie_list/pkg/server"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"log"
	"net"
)

func main() {
	port := ":50051"
	network := "tcp"
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := &server.Server{Players: map[uuid.UUID]string{}}
	protobuf.RegisterPlayerServiceServer(grpcServer, srv)

	log.Println("Starting server on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
