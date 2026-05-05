package service

import (
    "taskmanager/internal/model"
    "taskmanager/internal/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(u repository.UserRepository) *UserService {
    return &UserService{userRepo: u}
}

func (s *UserService) CreateUser(name string) error {
    user := &model.User{
        Name: name,
    }

    return s.userRepo.Create(user)
}

func (s *UserService) GetAllUsers() []*model.User {
    return s.userRepo.GetAll()
}