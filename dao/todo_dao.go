package dao

import (
	"context"
	"time"

	"todo-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoDAO struct {
	Collection *mongo.Collection
}

func NewTodoDAO(client *mongo.Client) *TodoDAO {
	collection := client.Database("todoDB").Collection("todos")
	return &TodoDAO{Collection: collection}
}

func (dao *TodoDAO) CreateTodo(todo model.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := dao.Collection.InsertOne(ctx, todo)
	return err
}

func (dao *TodoDAO) GetTodos() ([]model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todos []model.Todo
	cursor, err := dao.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
