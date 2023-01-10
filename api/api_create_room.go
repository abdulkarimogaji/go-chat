package api

import (
	"context"

	"github.com/abdulkarimogaji/go-chat/db"
	"github.com/abdulkarimogaji/go-chat/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {

	createdRoom, err := s.DbStore.CreateRoom(&db.Room{
		UserID:      int(req.GetUserId()),
		OtherUserId: int(req.GetOtherUserId()),
		RoomName:    req.GetRoomName(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create room %s", err)
	}

	return &pb.CreateRoomResponse{
		Room: &pb.Room{
			Id:          int32(createdRoom.Id),
			UserId:      int32(createdRoom.UserID),
			OtherUserId: int32(createdRoom.OtherUserId),
			RoomName:    createdRoom.RoomName,
			LastChatId:  int32(createdRoom.LastChatID),
			Status:      pb.RoomStatus(createdRoom.Status),
			CreatedAt:   timestamppb.New(createdRoom.CreatedAt),
			UpdatedAt:   timestamppb.New(createdRoom.UpdatedAt),
		},
	}, nil
}
