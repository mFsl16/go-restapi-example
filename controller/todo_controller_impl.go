package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/service"
	"github.com/mFsl16/go-restapi-example/utils"
	log "github.com/sirupsen/logrus"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: service,
	}
}

func (controller *TodoControllerImpl) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Header().Add("Content-Type", "application/json")
	message := "hello"
	encoder := json.NewEncoder(w)
	err := encoder.Encode(message)
	if err != nil {
		panic(err)
	}
}

func (controller *TodoControllerImpl) CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoRequest := model.Todo{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(&todoRequest)
	utils.PanicIfErr(err)

	log.Info("Receive request", todoRequest)

	todoCreated := controller.TodoService.CreateTodo(r.Context(), todoRequest)
	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   todoCreated,
	}

	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	errEncode := e.Encode(response)
	utils.PanicIfErr(errEncode)
}

func (controller *TodoControllerImpl) UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, errParse := strconv.Atoi(params.ByName("id"))
	utils.PanicIfErr(errParse)
	todoRequest := model.Todo{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(&todoRequest)
	utils.PanicIfErr(err)

	fmt.Println("======================================================================================")
	log.Info("Receive request [", todoRequest.Task, todoRequest.IsComplete, "][", id, "]")

	todoUpdated := controller.TodoService.UpdateTodo(r.Context(), id, todoRequest)

	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   todoUpdated,
	}

	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	errEncode := e.Encode(response)
	utils.PanicIfErr(errEncode)

	log.Info("Request completed [", response, "]")

}

func (controller *TodoControllerImpl) DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	utils.PanicIfErr(err)

	fmt.Println("======================================================================================")
	log.Info("Receive request [id: ", id, "]")

	deletedTodo := controller.TodoService.DeleteTodo(r.Context(), id)

	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   deletedTodo,
	}

	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	errEncode := e.Encode(response)
	utils.PanicIfErr(errEncode)

	log.Info("Request completed [", response, "]")
}

func (controller *TodoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id"))
	utils.PanicIfErr(err)

	fmt.Println("======================================================================================")
	log.Info("Receive request [id: ", id, "]")

	todoResult := controller.TodoService.FindTodoById(r.Context(), id)
	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   todoResult,
	}

	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	errEncode := e.Encode(response)
	utils.PanicIfErr(errEncode)

	log.Info("Request completed [", response, "]")

}

func (controller *TodoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Println("======================================================================================")
	log.Info("Receive request")

	allTodo := controller.TodoService.FindAllTodo(r.Context())
	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   allTodo,
	}

	w.Header().Add("Content-Type", "application/json")
	e := json.NewEncoder(w)
	errEncode := e.Encode(response)
	utils.PanicIfErr(errEncode)

	log.Info("Request completed [", response, "]")
}
