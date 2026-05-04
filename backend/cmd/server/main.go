package main

import (
    "net/http"
    "taskmanager/internal/handler"
    "taskmanager/internal/repository"
    "taskmanager/internal/service"
)

func main() {
    userRepo := repository.NewInMemoryUserRepo()
    taskRepo := repository.NewInMemoryTaskRepo()

    userService := service.NewUserService(userRepo)
    taskService := service.NewTaskService(taskRepo, userRepo)

    userHandler := handler.NewUserHandler(userService)
    taskHandler := handler.NewTaskHandler(taskService)

    http.HandleFunc("/users", userHandler.CreateUser)
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        taskHandler.GetTasks(w, r)
    case http.MethodPost:
        taskHandler.CreateTask(w, r)
    }
})

    http.HandleFunc("/tasks/done", taskHandler.MarkDone)

    http.ListenAndServe(":8080", nil)
}