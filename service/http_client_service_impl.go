package service

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/mFsl16/go-restapi-example/model"
	"github.com/mFsl16/go-restapi-example/utils"
)

type HttpClientServiceImpl struct {
	Client *resty.Client
}

func NewHttpClientService(client *resty.Client) HttpClientService {
	return &HttpClientServiceImpl{
		Client: client,
	}
}

func (service *HttpClientServiceImpl) GetPostById(ctx context.Context, id int) model.Post {

	response, err := service.Client.R().EnableTrace().Get("https://jsonplaceholder.typicode.com/todos/1")
	utils.PanicIfErr(err)
	post := model.Post{}
	errParse := json.Unmarshal(response.Body(), &post)
	utils.PanicIfErr(errParse)

	return post
}
