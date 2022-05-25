package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/iamtonydev/user-service-api/internal/app/api/user_v1"
	"github.com/iamtonydev/user-service-api/internal/app/repository"
	"github.com/iamtonydev/user-service-api/internal/app/service/user"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const (
	host       = "localhost"
	dbPort     = "54321"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "user_service_api"
	sslMode    = "disable"
	grpcPort   = ":50051"
	httpPort   = ":8000"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startGRPC())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startHTTP())
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to mapping port: %s", err.Error())
	}
	defer list.Close()

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return fmt.Errorf("failed to open connection with db")
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
	)
	desc.RegisterUserV1Server(s, user_v1.NewUserV1(userService))

	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to server: %s", err.Error())
	}

	return nil
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

	err := desc.RegisterUserV1HandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}
