package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/fabioods/fc_hexagonal_arch/adapters/dto"
	"github.com/fabioods/fc_hexagonal_arch/application"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("OPTIONS", "GET")
	router.Handle("/product", n.With(negroni.Wrap(createProduct(service)))).Methods("OPTIONS", "POST")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
