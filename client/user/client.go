package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "workbench-user/proto"
)

func main() {
	service := micro.NewService(micro.Name("workbench.cli.user"))

	service.Init()

	client := pb.NewUserServiceClient("workbench.service.user", service.Client())

	// get user list
	// func() {
	// 	req := &pb.ListRequest{
	// 		LastId: 0,
	// 		Offset: 5,
	// 		Limit:  10,
	// 		Team:   1,
	// 		Group:  2,
	// 	}

	// 	resp, err := client.List(context.TODO(), req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(resp)
	// }()

	// register
	req := &pb.RegisterRequest{
		Email:    "muxi@304.com",
		Name:     "muxi",
		Password: "muxi",
	}

	resp, err := client.Register(context.TODO(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// login
	// 	req := &pb.LoginRequest{
	// 		OauthCode: "NKPSPF0IOWWZSACLZ0OKAQ",
	// 	}

	// 	resp, err := client.Login(context.TODO(), req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(resp)
}
