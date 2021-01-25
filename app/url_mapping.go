package app

import (
	"github.com/yepack/testItem-api/controllers"
	"net/http"
)

func mapUrls(){
	router.HandleFunc("/items/", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/:?items", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}