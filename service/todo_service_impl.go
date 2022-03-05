package service

import (
	"context"
	"database/sql"

	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Mysql          *sql.DB
}

func NewTodoService(repository repository.TodoRepository, sql *sql.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: repository,
		Mysql:          sql,
	}
}

func (service *TodoServiceImpl) CreateTodo(ctx context.Context, sql *sql.DB, todo model.Todo) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) UpdateTodo(ctx context.Context, sql *sql.DB, id int, todo model.Todo) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) DeleteTodo(ctx context.Context, sql *sql.DB, id int) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) FindTodoById(ctx context.Context, sql *sql.DB, id int) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) FindAllTodo(ctx context.Context, sql *sql.DB) []model.Todo {
	return []model.Todo{}
}
