package handler

import (
	"context"

	"github.com/halhal23/strategy-user/domain"
	"github.com/halhal23/strategy-user/pkg/pb"
	"github.com/halhal23/strategy-user/usecase"
)

type UserHandler struct {
	UserInteractor *usecase.UserInteractor
}

func NewUserHandler(interactor *usecase.UserInteractor) *UserHandler {
	return &UserHandler{
		UserInteractor: interactor,
	}
}

func (handler *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	input := req.GetUserInput()
	user := domain.UserModel{
		Name: input.Name,
		Email: input.Email,
	}
	err := handler.UserInteractor.Add(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: 200}, nil
}

func (handler *UserHandler) ReadUser(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	id := req.GetId()
	user, err := handler.UserInteractor.UserById(id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadUserResponse{
		User: &pb.User{
			Name: user.Name,
			Email: user.Email,
		},
	}, nil
}

func (handler *UserHandler) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := handler.UserInteractor.Users()
	if err != nil {
		return nil, err
	}
	var res_users []*pb.User
	for _, user := range users {
		res_users = append(res_users, &pb.User{
			Id: user.ID,
			Name: user.Name,
			Email: user.Email,
		})
	}
	return &pb.ListUserResponse{
		User: res_users,	
	}, nil
}