package main

import (
	"github.com/iamtonydev/user-service-api/internal/app/api/user_v1"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterUserV1Server(s, user_v1.NewUserV1())

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to server: %s", err.Error())
	}
}
