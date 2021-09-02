package database

import (
	"fmt"
	"strategy-user/domain"
)

type UserRepo struct {
	SqlHandler
}

func (repo *UserRepo) Store(user domain.UserModel) (int64, error) {
	result, err := repo.Execute("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		fmt.Println("in interface/database/user_repository store method failed")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("in interface/database/user_repository LastInsertId method failed")
		return 0, err
	}
	return id, err
}

func (repo *UserRepo) FindById(indentifier int64) (user domain.UserModel, err error) {
	row, err := repo.Query("SELECT * FROM users WHERE id = ?", indentifier)
	if err != nil {
		fmt.Println("in interface/database/user_repository find by id method failed")
		return 
	}
	defer row.Close()
	row.Next()
	var id int64
	var name string
	var email string
	row.Scan(&id, &name, &email)
	user.ID = id
	user.Name = name 
	user.Email = email 
	return
}

func (repo *UserRepo) FindAll() (users []domain.UserModel, err error) {
	row, err := repo.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("in interface/database/user_repository find all method failed")
		return 
	}
	defer row.Close()
	for row.Next() {
		user := domain.UserModel{}
		row.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}
	return
}