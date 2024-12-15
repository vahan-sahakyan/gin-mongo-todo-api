package service

import (
	"todo-api/dao"
	"todo-api/model"
)

type TodoService struct {
	DAO *dao.TodoDAO
}

func (s *TodoService) CreateTodo(todo model.Todo) error {
	return s.DAO.CreateTodo(todo)
}

func (s *TodoService) GetTodos() ([]model.Todo, error) {
	return s.DAO.GetTodos()
}
