package controllers

import (
	"encoding/json"
	"github.com/yepack/testItem-api/domain/items"
	"github.com/yepack/testItem-api/services"
	"github.com/yepack/testItem-api/utils/http_utils"
	"github.com/yepack/testOauth_go/oauth"
	"github.com/yepack/testUtils-api/rest_errors"
	"io/ioutil"
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
	//Validating all the parameters  that we need to send this request to service
	if err := oauth.AuthenticateRequest(r); err != nil {
		//http_utils.RespondError(w, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		restErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}
	//end

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, restErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
