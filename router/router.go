package router

import (
	"github.com/gorilla/mux"
	"products_management/domain"
	"net/http"
)

//NewRouter for create router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	useCase := domain.GetProducts()

	//HandleFunc
	router.HandleFunc("/products", useCase.Get).Methods(http.MethodGet)
	router.HandleFunc("/products", useCase.Add).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", useCase.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/products/{id}", useCase.Edit).Methods(http.MethodPatch)
	router.HandleFunc("/products/{id}", useCase.GetDetail).Methods(http.MethodGet)

	//insert product category
	router.HandleFunc("/category", useCase.AddProductCategory).Methods(http.MethodPost)
	router.HandleFunc("/category", useCase.GetProductCategories).Methods(http.MethodGet)

	return router
}
