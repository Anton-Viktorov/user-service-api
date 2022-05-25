package main

import (
	"context"
	"fmt"
	"log"

	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithInsecure()) // nolint: static check
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewUserV1Client(con)

	// add user
	res, err := client.AddUser(ctx, &desc.AddUserRequest{
		Name:  "Admin",
		Age:   100500,
		Email: "admin@no.no",
	})
	if err != nil {
		log.Fatalf("failed to add user %s", err.Error())
	}

	fmt.Println("=== add user ===")
	fmt.Printf("user id: %d\n", res.GetResult().GetId())

	//get user
	user, err := client.GetUser(ctx, &desc.GetUserRequest{Id: 8})
	if err != nil {
		log.Fatalf("failed to get user %s", err.Error())
	}

	fmt.Println("=== user info ===")
	fmt.Printf("user: %s\n", user.GetResult())

	// multi add user
	reqData := []*desc.MultiAddUserRequest_User{
		{
			Name:  "Admin",
			Age:   100500,
			Email: "admin@no.no",
		},
		{
			Name:  "Admin2",
			Age:   100501,
			Email: "admin2@no.no",
		},
	}
	users, err := client.MultiAddUser(ctx, &desc.MultiAddUserRequest{Users: reqData})
	if err != nil {
		log.Fatalf("failed to multi add users %s", err.Error())
	}

	fmt.Println("=== users count ===")
	fmt.Printf("users added: %v\n", users.GetResult().GetCount())

	// list users
	listUsers, err := client.ListUser(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to get list users %s", err.Error())
	}

	fmt.Println("=== list users ===")
	fmt.Printf("users info: %s\n", listUsers.GetResult())

	// delete user
	_, err = client.RemoveUser(ctx, &desc.RemoveUserRequest{Id: 17})
	if err != nil {
		log.Fatalf("failed to remove user %s", err.Error())
	}

	// update user
	_, err = client.UpdateUser(ctx, &desc.UpdateUserRequest{
		Id:    20,
		Name:  "NewAdmin",
		Age:   100500,
		Email: "newadmin@no.no",
	})
	if err != nil {
		log.Fatalf("failed to update user %s", err.Error())
	}
}
