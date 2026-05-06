package repository

import (
    "errors"
    "fmt"
    "taskmanager/internal/model"
)

type InMemoryTaskRepo struct {
    tasks map[string]*model.Task
    nextID int
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
    return &InMemoryTaskRepo{
        tasks: make(map[string]*model.Task),
        nextID: 1,
    }
}

func (r *InMemoryTaskRepo) Create(task *model.Task) error {
    task.ID = fmt.Sprintf("%d", r.nextID)
    r.tasks[task.ID] = task
    r.nextID++
    return nil
}

func (r *InMemoryTaskRepo) GetByUser(userID string) ([]model.Task, error) {
    var result []model.Task
    for _, t := range r.tasks {
        if t.UserID == userID {
            result = append(result, *t)
        }
    }
    return result, nil
}

func (r *InMemoryTaskRepo) MarkDone(taskID int) error {
    task, ok := r.tasks[fmt.Sprintf("%d", taskID)]
    if !ok {
        return errors.New("task not found")
    }
    task.Done = true
    return nil
}