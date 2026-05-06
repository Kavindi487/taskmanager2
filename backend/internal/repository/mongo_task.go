package repository

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"

    "taskmanager/internal/db"
    "taskmanager/internal/model"
)

type MongoTaskRepo struct {
    col *mongo.Collection
}

func NewMongoTaskRepo() *MongoTaskRepo {
    return &MongoTaskRepo{col: db.Database.Collection("tasks")}
}

func (r *MongoTaskRepo) Create(task *model.Task) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := r.col.InsertOne(ctx, task)
    if err != nil {
        return err
    }
    task.ID = result.InsertedID.(string)
    return nil
}

func (r *MongoTaskRepo) GetByUser(userID string) ([]model.Task, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cursor, err := r.col.Find(ctx, bson.M{"user_id": userID})
    if err != nil {
        return nil, err
    }
    var tasks []model.Task
    if err := cursor.All(ctx, &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

func (r *MongoTaskRepo) MarkDone(taskID string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := r.col.UpdateOne(ctx,
        bson.M{"_id": taskID},
        bson.M{"$set": bson.M{"done": true}},
    )
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("task not found")
    }
    return nil
}