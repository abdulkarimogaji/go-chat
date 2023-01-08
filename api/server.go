package api

import (
	"github.com/abdulkarimogaji/go-chat/pb"
)

type Server struct {
	pb.UnimplementedGoChatServer
}

func NewServer() *Server {
	return &Server{}
}
