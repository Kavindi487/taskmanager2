package model

type Task struct {
    ID     string `bson:"_id,omitempty" json:"id"`
    Title  string `bson:"title"         json:"title"`
    Done   bool   `bson:"done"          json:"done"`
    UserID string `bson:"user_id"       json:"userId"`
}