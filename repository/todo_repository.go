package repository

import (
	"context"
	"database/sql"

	"github.com/mFsl16/go-restapi-example/model"
)

type TodoRepository interface {
	InsertTodo(ctx context.Context, db *sql.DB, todo model.Todo) (model.Todo, error)
	UpdateTodo(ctx context.Context, db *sql.DB, id int, todo model.Todo) (model.Todo, error)
	DeleteTodo(ctx context.Context, db *sql.DB, id int) bool
	FindById(ctx context.Context, db *sql.DB, id int) (model.Todo, error)
	FindAll(ctx context.Context, db *sql.DB) []model.Todo
}
