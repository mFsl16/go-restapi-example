package service

import (
	"context"

	"github.com/mFsl16/go-restapi-example/model"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo model.Todo) model.Todo
	UpdateTodo(ctx context.Context, id int, todo model.Todo) model.Todo
	DeleteTodo(ctx context.Context, id int) model.Todo
	FindTodoById(ctx context.Context, id int) model.Todo
	FindAllTodo(ctx context.Context) []model.Todo
}
