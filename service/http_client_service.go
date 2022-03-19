package service

import (
	"context"

	"github.com/mFsl16/go-restapi-example/model"
)

type HttpClientService interface {
	GetPostById(ctx context.Context, id int) model.Post
}
