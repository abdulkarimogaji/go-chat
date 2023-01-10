package api

import (
	"context"
	"time"

	"github.com/abdulkarimogaji/go-chat/db"
	"github.com/abdulkarimogaji/go-chat/pb"
	"github.com/abdulkarimogaji/go-chat/token"
	"github.com/abdulkarimogaji/go-chat/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	createdUser, err := s.DbStore.CreateUser(&db.User{
		Username: req.GetUsername(),
		Fullname: req.GetFullname(),
		Email:    req.GetEmail(),
		Password: hashedPassword,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user %s", err)
	}

	// generate token
	tokenMaker, err := token.NewJwtMaker("abdul")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating token maker %s", err)
	}

	tokenStr, err := tokenMaker.CreateToken(int(createdUser.Id), time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating token %s", err)
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
		AccessToken: tokenStr,
	}, nil
}
