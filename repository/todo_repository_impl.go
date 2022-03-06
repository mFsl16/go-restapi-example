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

func (repository *TodoRepositoryImpl) InsertTodo(ctx context.Context, db *sql.DB, todo model.Todo) (model.Todo, error) {
	sql := "insert into todo (task, is_complete, created_at) values(?, ?, CURDATE())"
	r, err := db.ExecContext(ctx, sql, todo.Task, todo.IsComplete)
	utils.PanicIfErr(err)
	id, err := r.LastInsertId()
	if err != nil {
		return model.Todo{}, err
	}
	todo.Id = int(id)
	return todo, nil
}

func (repository *TodoRepositoryImpl) UpdateTodo(ctx context.Context, db *sql.DB, id int, todo model.Todo) (model.Todo, error) {
	sql := "update todo set task = ?, is_complete = ?, update_at = CURDATE() where id = ?"
	_, err := db.ExecContext(ctx, sql, todo.Task, todo.IsComplete, id)
	if err != nil {
		utils.PanicIfErr(err)
		return model.Todo{}, err
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) DeleteTodo(ctx context.Context, db *sql.DB, id int) bool {

	sql := "delete from todo where id = ?"
	_, err := db.ExecContext(context.Background(), sql, id)
	utils.PanicIfErr(err)

	return err == nil
}

func (repository *TodoRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int) (model.Todo, error) {

	sql := "select id, task, is_complete, created_at, update_at from todo where id = ?"
	r, err := db.QueryContext(context.Background(), sql, id)
	if err != nil {
		utils.PanicIfErr(err)
		return model.Todo{}, err
	}

	defer r.Close()

	todo := model.Todo{}

	if r.Next() {
		errQuery := r.Scan(&todo.Id, &todo.Task, &todo.IsComplete, &todo.CreateAt, &todo.UpdateAt)
		utils.PanicIfErr(errQuery)
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []model.Todo {

	todos := []model.Todo{}

	sql := "select id, task, is_complete, created_at, update_at from todo"

	r, err := db.QueryContext(ctx, sql)
	utils.PanicIfErr(err)
	defer r.Close()

	for r.Next() {
		todo := model.Todo{}
		errQuery := r.Scan(&todo.Id, &todo.Task, &todo.IsComplete, &todo.CreateAt, &todo.UpdateAt)
		utils.PanicIfErr(errQuery)
		todos = append(todos, todo)
	}

	return todos
}
