package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/service"
	"github.com/mFsl16/go-restapi-example/utils"
	log "github.com/sirupsen/logrus"
)

type PostControllerImpl struct {
	Service service.HttpClientService
}

func NewPostController(service service.HttpClientService) PostController {

	return &PostControllerImpl{
		Service: service,
	}
}

func (controller *PostControllerImpl) FindPostById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id"))
	utils.PanicIfErr(err)

	log.Info("Receive request [id: ", id, "]")

	post := controller.Service.GetPostById(r.Context(), id)
	response := model.CommonResponse{
		Code:   200,
		Status: "success",
		Data:   post,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err2 := encoder.Encode(response)
	utils.PanicIfErr(err2)
}
