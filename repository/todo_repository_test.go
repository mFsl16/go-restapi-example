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
	result := repository.InsertTodo(context.Background(), db.Mysql, todo)

	assert.Equal(t, todo.Task, result.Task)
}
