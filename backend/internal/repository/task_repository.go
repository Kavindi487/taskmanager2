package repository

import "taskmanager/internal/model"

type TaskRepository interface {
    Create(task *model.Task) error
    GetByUser(userID int) ([]model.Task, error)
    MarkDone(taskID int) error
}