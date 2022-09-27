package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutesForStories(router *mux.Router) {
	// DAR PERMISOS CORS
	enableCORS(router)

	router.HandleFunc("/stories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stories, err := getImages()
		if err == nil {
			respondWithSuccess(stories, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	//esto es para onteber un solo id
	router.HandleFunc("/stories/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		stories, err := getImageById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(stories, w)
		}
	}).Methods(http.MethodGet)

	//esto es para crear una categoria
	router.HandleFunc("/stories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

		var stories Stories
		err := json.NewDecoder(r.Body).Decode(&stories)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createImages(stories)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/stories/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)

		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateImages(id)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	// router.HandleFunc("/categorias/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json") //apra formatear en tipo json

	// 	idAsString := mux.Vars(r)["id"]
	// 	id, err := stringToInt64(idAsString)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 		// We return, so we stop the function flow
	// 		return
	// 	}
	// 	err = deleteImages(id)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 	} else {
	// 		respondWithSuccess(true, w)
	// 	}
	// }).Methods(http.MethodDelete)
}
