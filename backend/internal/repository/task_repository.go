package repository

import "taskmanager/internal/model"

type TaskRepository interface {
    Create(task *model.Task) error
    GetByUser(userID string) ([]model.Task, error)
    MarkDone(taskID string) error
}