package main

import (
    "net/http"
    "taskmanager/internal/handler"
    "taskmanager/internal/repository"
    "taskmanager/internal/service"
)

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(200)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    userRepo := repository.NewInMemoryUserRepo()
    taskRepo := repository.NewInMemoryTaskRepo()

    userService := service.NewUserService(userRepo)
    taskService := service.NewTaskService(taskRepo, userRepo)

    userHandler := handler.NewUserHandler(userService)
    taskHandler := handler.NewTaskHandler(taskService)

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            userHandler.GetUsers(w, r)
        case http.MethodPost:
            userHandler.CreateUser(w, r)
        }
    })

    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            taskHandler.GetTasks(w, r)
        case http.MethodPost:
            taskHandler.CreateTask(w, r)
        }
    })

    http.HandleFunc("/tasks/done", taskHandler.MarkDone)

    http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
}