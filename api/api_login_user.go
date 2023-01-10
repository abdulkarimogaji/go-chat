package api

import (
	"context"
	"time"

	"github.com/abdulkarimogaji/go-chat/pb"
	"github.com/abdulkarimogaji/go-chat/token"
	"github.com/abdulkarimogaji/go-chat/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// get user by username
	foundUser, err := s.DbStore.GetUserByUsername(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found %s", err)
	}

	// compare password
	err = utils.CheckPassword(req.GetPassword(), foundUser.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect username or password %s", err)
	}

	// generate token
	tokenMaker, err := token.NewJwtMaker("abdul")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating token maker %s", err)
	}

	tokenStr, err := tokenMaker.CreateToken(int(foundUser.Id), time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating token %s", err)
	}

	return &pb.LoginResponse{
		User: &pb.User{
			Id:        int32(foundUser.Id),
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			Role:      pb.Role(foundUser.Role),
			Fullname:  foundUser.Fullname,
			CreatedAt: timestamppb.New(foundUser.CreatedAt),
			UpdatedAt: timestamppb.New(foundUser.UpdatedAt),
		},
		AccessToken: tokenStr,
	}, nil
}
