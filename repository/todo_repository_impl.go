package repository

import (
	"context"
	"database/sql"

	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/utils"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) InsertTodo(ctx context.Context, db *sql.DB, todo model.Todo) model.Todo {
	sql := "insert into todo (task, is_complete, create_at) values(?, ?, CURDATE())"
	r, err := db.ExecContext(ctx, sql, todo.Task, todo.IsComplete)
	utils.PanicIfErr(err)
	id, err := r.LastInsertId()
	utils.PanicIfErr(err)
	todo.Id = int(id)
	return todo
}

func (repository *TodoRepositoryImpl) UpdateTodo(ctx context.Context, db *sql.DB, id int, todo model.Todo) model.Todo {
	return todo
}

func (repository *TodoRepositoryImpl) DeleteTodo(ctx context.Context, db *sql.DB, id int) bool {

	return true
}

func (repository *TodoRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int) model.Todo {

	return model.Todo{}
}

func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []model.Todo {

	return []model.Todo{}
}
