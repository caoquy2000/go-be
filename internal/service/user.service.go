package service

import "go-be/internal/repo"

type UserService struct {
	ur *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		ur: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.ur.GetUserInfo()
}