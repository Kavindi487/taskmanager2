package repository

import (
    "errors"
    "taskmanager/internal/model"
)

type InMemoryUserRepo struct {
    users map[int]*model.User
    nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{
        users: make(map[int]*model.User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepo) Create(user *model.User) error {
    user.ID = r.nextID
    r.users[user.ID] = user
    r.nextID++
    return nil
}

func (r *InMemoryUserRepo) GetByID(id int) (*model.User, error) {
    user, ok := r.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return user, nil
}