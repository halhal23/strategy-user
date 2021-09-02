package controller

import (
	"strategy-user/domain"
	"strategy-user/interface/database"
	"strategy-user/usecase"
	"strconv"
)

type UserController struct {
	UserInteractor *usecase.UserInteractor
}

func NewUserController(sqlHandler *database.SqlHandler) *UserController {
	return &UserController{
		UserInteractor: &usecase.UserInteractor{
			UserRepo: &database.UserRepo{
				SqlHandler: *sqlHandler,
			},
		},
	}
}

func (controller *UserController) Index(c Context) {
	users, err := controller.UserInteractor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	user, err := controller.UserInteractor.UserById(int64(id))
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) Create(c Context) {
	var user domain.UserModel
	c.Bind(&user)
	err := controller.UserInteractor.Add(user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, nil)
}
