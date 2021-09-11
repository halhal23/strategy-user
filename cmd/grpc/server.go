package main

import (
	"fmt"
	"net"

	"github.com/halhal23/strategy-user/infrastructure"
	"github.com/halhal23/strategy-user/interface/database"
	"github.com/halhal23/strategy-user/interface/grpc/handler"
	"github.com/halhal23/strategy-user/pkg/pb"
	"github.com/halhal23/strategy-user/usecase"

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