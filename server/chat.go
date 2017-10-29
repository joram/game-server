package server

import (
	"context"

	"errors"
	"github.com/joram/game-server/pb"
)

func (s *GameServer) RegisterToChannel(request *pb.RegisterToChannelRequest, server pb.GameServer_RegisterToChannelServer) error {
	return errors.New("unimplemented")
}

func (s *GameServer) SendToChannel(ctx context.Context, request *pb.SendToChannelRequest) (*pb.SendToChannelResponse, error) {
	return nil, errors.New("unimplemented")
}
