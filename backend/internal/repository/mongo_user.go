package repository

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"

    "taskmanager/internal/db"
    "taskmanager/internal/model"
)

type MongoUserRepo struct {
    col *mongo.Collection
}

func NewMongoUserRepo() *MongoUserRepo {
    return &MongoUserRepo{col: db.Database.Collection("users")}
}

func (r *MongoUserRepo) Create(user *model.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := r.col.InsertOne(ctx, user)
    if err != nil {
        return err
    }
    user.ID = result.InsertedID.(string)
    return nil
}

func (r *MongoUserRepo) GetByID(id string) (*model.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    var user model.User
    err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if errors.Is(err, mongo.ErrNoDocuments) {
        return nil, errors.New("user not found")
    }
    return &user, err
}

func (r *MongoUserRepo) GetAll() []*model.User {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cursor, err := r.col.Find(ctx, bson.D{}, options.Find())
    if err != nil {
        return nil
    }
    var users []*model.User
    cursor.All(ctx, &users)
    return users
}