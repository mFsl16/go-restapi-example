package service

import (
	"context"
	"testing"

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
		CreateAt:   "2022-2-6",
	}
	context := context.Background()
	db := repository.NewDB()
	mysql := db.Mysql

	mockrepository.EXPECT().InsertTodo(context, mysql, todo).Return(todo, nil)

	service := NewTodoService(mockrepository, db)
	todoResult := service.CreateTodo(context, todo)
	assert.NotNil(t, todoResult.Id)
}
