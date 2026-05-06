package handler

import (
    "encoding/json"
    "net/http"
    "taskmanager/internal/service"
)

type TaskHandler struct {
    service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
    return &TaskHandler{service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    title  := r.FormValue("title")
    userID := r.FormValue("user")

    if title == "" {
        http.Error(w, "title is required", 400)
        return
    }
    if userID == "" {
        http.Error(w, "user id is required", 400)
        return
    }

    if err := h.service.CreateTask(title, userID); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    w.Write([]byte("Task created"))
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user")
    if userID == "" {
        http.Error(w, "user id required", 400)
        return
    }
    tasks, err := h.service.GetTasksByUser(userID)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) MarkDone(w http.ResponseWriter, r *http.Request) {
    taskID := r.URL.Query().Get("id")
    if taskID == "" {
        http.Error(w, "task id required", 400)
        return
    }
    if err := h.service.MarkDone(taskID); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    w.Write([]byte("Task marked as done"))
}