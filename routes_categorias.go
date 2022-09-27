package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// LO QUE PUSE DE API CATEGORIAS
func setupRoutesForCategorias(router *mux.Router) {
	// DAR PERMISOS CORS
	enableCORS(router)

	router.HandleFunc("/categorias/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		categorias, err := getCategorias()
		if err == nil {
			respondWithSuccess(categorias, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	//esto es para onteber un solo id
	router.HandleFunc("/categorias/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		categoria, err := getCategoriaById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(categoria, w)
		}
	}).Methods(http.MethodGet)

	//esto es para crear una categoria
	router.HandleFunc("/categorias/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json
		var categorias Categorias
		err := json.NewDecoder(r.Body).Decode(&categorias)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createCategorias(categorias)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/categorias/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateCategorias(id)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/categoria/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := activar(id)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/categorias/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		err = deleteCategorias(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}
