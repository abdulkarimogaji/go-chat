package api

import (
	"context"
	"time"

	"github.com/abdulkarimogaji/go-chat/db"
	"github.com/abdulkarimogaji/go-chat/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword := req.Password + "hashed"

	createdUser, err := s.DbStore.CreateUser(&db.User{
		Username: req.Username,
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user %s", err)
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Username:  createdUser.Username,
			Fullname:  createdUser.Fullname,
			Email:     createdUser.Email,
			Role:      pb.Role_MEMBER,
			CreatedAt: timestamppb.New(createdUser.CreatedAt),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}
