package main

import (
	"log"
	"net"

	"github.com/abdulkarimogaji/go-chat/api"
	"github.com/abdulkarimogaji/go-chat/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var PORT = ":8080"

func main() {
	err := startServer(PORT)
	if err != nil {
		log.Fatal("failed to start server", err)
	}

}

func startServer(port string) error {
	grpcServer := grpc.NewServer()
	server := api.NewServer()
	pb.RegisterGoChatServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", port)

	if err != nil {
		return err
	}
	log.Println("Server listening on port", port)
	return grpcServer.Serve(listener)
}
