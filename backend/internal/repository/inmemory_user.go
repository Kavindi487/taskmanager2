package repository

import (
    "errors"
    "sync"
    "taskmanager/internal/model"
)

type InMemoryUserRepo struct {
    mu     sync.RWMutex
    users  map[int]*model.User
    nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{
        users: make(map[int]*model.User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepo) Create(user *model.User) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    user.ID = r.nextID
    r.users[user.ID] = user
    r.nextID++
    return nil
}

func (r *InMemoryUserRepo) GetAll() []*model.User {
    r.mu.RLock()
    defer r.mu.RUnlock()
    users := make([]*model.User, 0, len(r.users))
    for _, u := range r.users {
        users = append(users, u)
    }
    return users
}

func (r *InMemoryUserRepo) GetByID(id int) (*model.User, error) {
    user, ok := r.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return user, nil
}