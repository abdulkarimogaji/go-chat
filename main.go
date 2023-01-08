package main

import (
	"log"
	"net"

	"github.com/abdulkarimogaji/go-chat/api"
	"github.com/abdulkarimogaji/go-chat/db"
	"github.com/abdulkarimogaji/go-chat/pb"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var PORT = ":8080"
var DB_URI = "root:password@tcp(localhost:3000)/go_chat?parseTime=true"

func main() {
	store, err := db.ConnectDatabase(DB_URI)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = startServer(store, PORT)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}

func startServer(store *db.Store, port string) error {
	server := api.NewServer(store)

	grpcServer := grpc.NewServer()
	pb.RegisterGoChatServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", port)

	if err != nil {
		return err
	}
	log.Println("Server listening on port", port)
	return grpcServer.Serve(listener)
}
