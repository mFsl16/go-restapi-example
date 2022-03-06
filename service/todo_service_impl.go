package service

import (
	"context"

	"github.com/mFsl16/go-restapi-example/exception"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Repo           *repository.Database
}

func NewTodoService(todoRepository repository.TodoRepository, repo *repository.Database) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		Repo:           repo,
	}
}

func (service *TodoServiceImpl) CreateTodo(ctx context.Context, todo model.Todo) model.Todo {

	todo, err := service.TodoRepository.InsertTodo(ctx, service.Repo.Mysql, todo)
	if err != nil {
		panic(exception.NewCommonException(err.Error()))
	}
	return todo
}

func (service *TodoServiceImpl) UpdateTodo(ctx context.Context, id int, todo model.Todo) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) DeleteTodo(ctx context.Context, id int) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) FindTodoById(ctx context.Context, id int) model.Todo {
	return model.Todo{}
}

func (service *TodoServiceImpl) FindAllTodo(ctx context.Context) []model.Todo {
	return []model.Todo{}
}
