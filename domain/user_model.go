package domain

type UserModel struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func NewUserModel(id int64, name, email string) *UserModel {
	return &UserModel{ID: id, Name: name, Email: email}
}
