package controllers

import (
	"fmt"
	"github.com/yepack/testItem-api/domain/items"
	"github.com/yepack/testItem-api/services"
	"github.com/yepack/testOauth_go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error ti the caller.
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil{
		//TODO: Return error json to the client
	}
	fmt.Println(result)
	//TODO: Create item with HTTP status 201 - Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
