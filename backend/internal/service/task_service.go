package service

import (
    "taskmanager/internal/model"
    "taskmanager/internal/repository"
)

type TaskService struct {
    taskRepo repository.TaskRepository
    userRepo repository.UserRepository
}

func NewTaskService(t repository.TaskRepository, u repository.UserRepository) *TaskService {
    return &TaskService{
        taskRepo: t,
        userRepo: u,
    }
}

func (s *TaskService) CreateTask(title string, userID int) error {
    _, err := s.userRepo.GetByID(userID)
    if err != nil {
        return err
    }

    task := &model.Task{
        Title:  title,
        Done:   false,
        UserID: userID,
    }

    return s.taskRepo.Create(task)
}

func (s *TaskService) GetTasksByUser(userID int) ([]model.Task, error) {
    // optional: check if user exists
    _, err := s.userRepo.GetByID(userID)
    if err != nil {
        return nil, err
    }

    return s.taskRepo.GetByUser(userID)
}

func (s *TaskService) MarkDone(taskID int) error {
    return s.taskRepo.MarkDone(taskID)
}