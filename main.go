package main

import (
	"context"
	"log"

	"todo-api/controller"
	"todo-api/dao"
	"todo-api/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup dependencies
	todoDAO := dao.NewTodoDAO(client)
	todoService := &service.TodoService{DAO: todoDAO}
	todoController := &controller.TodoController{Service: todoService}

	// Setup Gin router
	router := gin.Default()
	todoController.RegisterRoutes(router)

	// Start the server
	log.Println("Starting server on port 8080...")
	router.Run(":8080")
}
