package server

type GameServer struct {
}

func New() (*GameServer, error) {
	return &GameServer{}, nil
}
