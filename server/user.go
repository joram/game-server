package server

import (
	"context"
	"fmt"

	"errors"
	"github.com/joram/game-server/pb"
)

func (s *GameServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, errors.New("unimplemented")
}

func (s *GameServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("'%v' logged in\n", request.Username)
	return &pb.LoginResponse{}, nil
}
