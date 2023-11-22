package service

import (
	"strconv"
	"youandus/pkg/users/model"
	user2 "youandus/pkg/users/model/user"
	"youandus/pkg/users/repository"
)

type UsersService struct {
	repository *repository.UsersRepository
}

func NewUsersService(repository *repository.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (s *UsersService) CreateUser(user *user2.UserRegister) (*model.BaseResponse, *model.BaseErrorResponse) {
	registerUser, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, &model.BaseErrorResponse{
			Message: "User can't created !",
			Error:   err.Error(),
		}
	}
	return registerUser, nil
}
func (s *UsersService) LoginUser(user *user2.UserLogin) (*model.BaseResponse, *model.BaseErrorResponse) {
	loginUser, err := s.repository.LoginUser(user)
	if err != nil {
		return nil, &model.BaseErrorResponse{
			Message: "User can't logined !",
			Error:   err.Error(),
		}
	}
	return loginUser, nil
}
func (s *UsersService) UpdateUser(id uint, user *user2.UserRegister) (*model.BaseResponse, *model.BaseErrorResponse) {
	updateUser, err := s.repository.UpdateUser(id, user)
	if err != nil {
		return nil, &model.BaseErrorResponse{
			Message: "User can't updated !",
			Error:   err.Error(),
		}
	}
	return updateUser, nil
}
func (s *UsersService) DeleteUser(id string) *model.BaseErrorResponse {
	intID, _ := strconv.Atoi(id)
	if err := s.repository.DeleteUser(uint(intID)); err != nil {
		return &model.BaseErrorResponse{
			Message: "User can't deleted !",
			Error:   err.Error(),
		}
	}
	return nil
}

func (s *UsersService) GetUsers() (*model.BaseResponse, *model.BaseErrorResponse) {
	_, err := s.repository.GetUsers()
	if err != nil {
		return nil, &model.BaseErrorResponse{
			Message: "Users can't get !",
			Error:   err.Error(),
		}
	}
	return &model.BaseResponse{
		Message: "ileride düzenleencek",
	}, nil
}
