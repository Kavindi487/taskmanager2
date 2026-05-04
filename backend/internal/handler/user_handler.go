package handler

import (
    "net/http"
    "taskmanager/internal/service"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")

    if name == "" {
        http.Error(w, "name is required", 400)
        return
    }

    err := h.service.CreateUser(name)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Write([]byte("User created"))
}