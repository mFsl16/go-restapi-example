package service

import (
	"context"

	"github.com/mFsl16/go-restapi-example/exception"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/repository"
	"github.com/mFsl16/go-restapi-example/utils"
	log "github.com/sirupsen/logrus"
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

	todoResult, err := service.TodoRepository.FindById(ctx, service.Repo.Mysql, id)
	if err != nil {
		panic(exception.NewCommonException(err.Error()))
	}

	log.Info("Todo Found", todoResult)

	todo.Id = todoResult.Id

	log.Info("updating todo start [", todo, "]")
	t, errUpdate := service.TodoRepository.UpdateTodo(ctx, service.Repo.Mysql, id, todo)
	if errUpdate != nil {
		panic(exception.NewCommonException(err.Error()))
	}
	log.Info("updating todo end [", t, "]")

	after, errAfter := service.TodoRepository.FindById(ctx, service.Repo.Mysql, id)
	utils.PanicIfErr(errAfter)

	return after
}

func (service *TodoServiceImpl) DeleteTodo(ctx context.Context, id int) model.Todo {

	todoResult, err := service.TodoRepository.FindById(ctx, service.Repo.Mysql, id)
	if err != nil {
		panic(exception.NewCommonException(err.Error()))
	}

	isSuccess := service.TodoRepository.DeleteTodo(ctx, service.Repo.Mysql, id)

	if !isSuccess {
		panic(exception.NewCommonException("failed to delete todo, todo not found"))
	}

	return todoResult
}

func (service *TodoServiceImpl) FindTodoById(ctx context.Context, id int) model.Todo {

	todoResult, err := service.TodoRepository.FindById(ctx, service.Repo.Mysql, id)
	if err != nil || todoResult.Id == 0 {
		panic(exception.NewCommonException("todo not found"))
	}

	return todoResult
}

func (service *TodoServiceImpl) FindAllTodo(ctx context.Context) []model.Todo {

	t := service.TodoRepository.FindAll(ctx, service.Repo.Mysql)
	return t
}
