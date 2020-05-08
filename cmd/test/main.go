package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/joram/game-server/pb"
)

func main() {
	fmt.Println("testing")
	serverAddr := "localhost:2303"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGameServerClient(conn)

	resp, err := client.Login(context.Background(), &pb.LoginRequest{Username: "test username"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp)
}
