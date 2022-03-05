package service

import (
	"context"
	"database/sql"

	"github.com/mFsl16/go-restapi-example/model"
)

type TodoService interface {
	CreateTodo(ctx context.Context, sql *sql.DB, todo model.Todo) model.Todo
	UpdateTodo(ctx context.Context, sql *sql.DB, id int, todo model.Todo) model.Todo
	DeleteTodo(ctx context.Context, sql *sql.DB, id int) model.Todo
	FindTodoById(ctx context.Context, sql *sql.DB, id int) model.Todo
	FindAllTodo(ctx context.Context, sql *sql.DB) []model.Todo
}
