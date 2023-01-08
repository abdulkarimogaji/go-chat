package api

import (
	"github.com/abdulkarimogaji/go-chat/db"
	"github.com/abdulkarimogaji/go-chat/pb"
)

type Server struct {
	pb.UnimplementedGoChatServer
	DbStore *db.Store
}

func NewServer(store *db.Store) *Server {
	return &Server{DbStore: store}
}
