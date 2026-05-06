package repository

import "taskmanager/internal/model"

type UserRepository interface {
    Create(user *model.User) error
    GetByID(id string) (*model.User, error)
    GetAll() []*model.User
}