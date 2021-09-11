package main

import (
	"fmt"
	"net"
	"strategy-user/infrastructure"
	"strategy-user/interface/database"
	"strategy-user/interface/grpc/handler"
	"strategy-user/pkg/pb"
	"strategy-user/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()

	interactor := usecase.NewUserInteractor(&database.UserRepo{infrastructure.NewSqlHandler()})
	handler := handler.NewUserHandler(interactor)

	pb.RegisterUserServiceServer(server, handler)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		fmt.Println(err)
	}
}