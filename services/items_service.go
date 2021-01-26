package services

import (
	"github.com/yepack/testItem-api/domain/items"
	"github.com/yepack/testUtils-api/rest_errors"
	"net/http"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}

	return &itemRequest, nil
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me",
		Error:   "not_implemented",
	}
}
