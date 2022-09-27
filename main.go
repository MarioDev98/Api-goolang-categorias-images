package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Conexion a la base de datos
	bd, err := getDB()
	if err != nil {
		log.Printf("Error al conextar en la base datos" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al establecer conexion a la base de datos, revisa .env  : " + err.Error())
			return
		}
	}

	//Definir las rutas
	router := mux.NewRouter()
	setupRoutesForStories(router)
	setupRoutesForCategorias(router)

	//se pueden poner mas condiguraciones
	// Puerto en donde se va ejecutar la base de datos
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// tiempo que se va tardar en responder al leer y escribir
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server funcionando en el puerto %s", port)
	log.Fatal(server.ListenAndServe())
}
