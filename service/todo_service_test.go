package service

import (
	"context"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang/mock/gomock"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/repository"
	"github.com/mFsl16/go-restapi-example/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	mockrepository := mocks.NewMockTodoRepository(gomock.NewController(t))
	todo := model.Todo{
		Id:         1,
		Task:       "Learn Golang",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}
	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mockrepository.EXPECT().InsertTodo(context, mysql, todo).Return(todo, nil)

	service := NewTodoService(mockrepository, db)
	todoResult := service.CreateTodo(context, todo)
	assert.NotNil(t, todoResult.Id)
}

func TestUpdateTodo(t *testing.T) {
	mockrepository := mocks.NewMockTodoRepository(gomock.NewController(t))
	todo := model.Todo{
		Id:         1,
		Task:       "Learn Golang",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}

	todoUpdate := model.Todo{
		Id:         1,
		Task:       "Updated Task",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}

	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mockrepository.EXPECT().FindById(context, mysql, 1).Return(todo, nil)
	mockrepository.EXPECT().UpdateTodo(context, mysql, 1, todo).Return(todoUpdate, nil)

	service := NewTodoService(mockrepository, db)

	todoResult := service.UpdateTodo(context, 1, todo)
	assert.Equal(t, todoUpdate.Task, todoResult.Task)
}

func TestDeleteTodo(t *testing.T) {
	mtr := mocks.NewMockTodoRepository(gomock.NewController(t))
	todo := model.Todo{
		Id:         1,
		Task:       "Learn Golang",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}

	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mtr.EXPECT().FindById(context, mysql, 1).Return(todo, nil)
	mtr.EXPECT().DeleteTodo(context, mysql, 1).Return(true)

	service := NewTodoService(mtr, db)

	todoResult := service.DeleteTodo(context, 1)

	assert.Equal(t, todo.Task, todoResult.Task)
}

func TestFindById(t *testing.T) {
	mtr := mocks.NewMockTodoRepository(gomock.NewController(t))

	todo := model.Todo{
		Id:         1,
		Task:       "Learn Golang",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}

	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mtr.EXPECT().FindById(context, mysql, 1).Return(todo, nil)

	service := NewTodoService(mtr, db)

	todoResult := service.FindTodoById(context, 1)

	assert.Equal(t, todo.Task, todoResult.Task)
}

func TestFindAllTodo(t *testing.T) {
	mtr := mocks.NewMockTodoRepository(gomock.NewController(t))
	todo := model.Todo{
		Id:         1,
		Task:       "Learn Golang",
		IsComplete: false,
		CreateAt:   time.Now().String(),
	}

	todos := []model.Todo{}

	todos = append(todos, todo)

	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mtr.EXPECT().FindAll(context, mysql).Return(todos)

	service := NewTodoService(mtr, db)

	todosResult := service.FindAllTodo(context)
	assert.True(t, len(todosResult) > 0)
}
