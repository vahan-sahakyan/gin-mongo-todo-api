package controller

import (
	"net/http"
	"todo-api/model"
	"todo-api/service"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Service *service.TodoService
}

func (c *TodoController) RegisterRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/todos")
	todoRoutes.POST("", c.CreateTodoHandler)
	todoRoutes.GET("", c.GetTodosHandler)
}

func (c *TodoController) CreateTodoHandler(ctx *gin.Context) {
	var todo model.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateTodo(todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Todo created successfully!"})
}

func (c *TodoController) GetTodosHandler(ctx *gin.Context) {
	todosData, err := c.Service.GetTodos()
	state := &State{
		Data:   todosData,
		Length: len(todosData),
	}

	if state.Length%2 == 1 {
		state.Message = "You have an ODD number of todos"
	} else {
		state.Message = "You have an EVEN number of todos"
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    state.Data,
		"length":  state.Length,
		"message": state.Message,
	})
}

type State struct {
	Data    []model.Todo
	Length  int
	Message string
}
