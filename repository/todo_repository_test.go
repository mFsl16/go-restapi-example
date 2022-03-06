package repository

import (
	"context"
	"testing"

	"github.com/mFsl16/go-restapi-example/model"
	"github.com/stretchr/testify/assert"
)

func TestInsertTodo(t *testing.T) {
	repository := NewTodoRepository()
	db := NewDB()

	todo := model.Todo{
		Task:       "Learn Golang",
		IsComplete: false,
	}
	result, err := repository.InsertTodo(context.Background(), db.Mysql, todo)

	assert.Nil(t, err)
	assert.Equal(t, todo.Task, result.Task)
}

func TestUpdateTodo(t *testing.T) {
	repository := NewTodoRepository()
	db := NewDB()

	todo := model.Todo{
		Id:         1,
		Task:       "Updated Task",
		IsComplete: true,
	}

	result, err := repository.UpdateTodo(context.Background(), db.Mysql, todo.Id, todo)

	assert.Nil(t, err)
	assert.Equal(t, todo.Task, result.Task)

}

func TestDeleteTodo(t *testing.T) {
	db := NewDB()
	repository := NewTodoRepository()

	result := repository.DeleteTodo(context.Background(), db.Mysql, 3)

	assert.True(t, result)
}

func TestFindByid(t *testing.T) {
	tr := NewTodoRepository()
	db := NewDB().Mysql

	todo, err := tr.FindById(context.Background(), db, 4)

	assert.Nil(t, err)
	assert.Equal(t, "Learn Golang", todo.Task)

}

func TestFindAll(t *testing.T) {
	tr := NewTodoRepository()
	db := NewDB().Mysql

	todos := tr.FindAll(context.Background(), db)

	assert.Len(t, todos, 1)

}
