package handler

import (
	"net/http"
	"strconv"
	"taskmanager/internal/service"
)

type TaskHandler struct {
    service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
    return &TaskHandler{service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")

    err := h.service.CreateTask(title, 1)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    w.Write([]byte("Task created"))
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("user")

    if userIDStr == "" {
        http.Error(w, "user id required", 400)
        return
    }

    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "invalid user id", 400)
        return
    }

    tasks, err := h.service.GetTasksByUser(userID)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    for _, t := range tasks {
        w.Write([]byte(t.Title + "\n"))
    }
}

func (h *TaskHandler) MarkDone(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    if idStr == "" {
        http.Error(w, "task id required", 400)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid task id", 400)
        return
    }

    err = h.service.MarkDone(id)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    w.Write([]byte("Task marked as done"))
}