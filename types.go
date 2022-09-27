package main

// modelo de datos de api images
// el primer campo es para ponerlo en controlador como variables globales
// el json: {campo} en campo tiene que ser escrito tal cual esta en la base de datos
type Stories struct {
	Id        int64  `json:"id"`
	Nombre    string `json:"nombre"`
	Categoria string `json:"categoria"`
	Ruta      string `json:"ruta"`
	Status    int64  `json:"status"`
}

// modelo de datos de api categorias
// el primer campo es para ponerlo en controlador como variables globales
// el json: {campo} en campo tiene que ser escrito tal cual esta en la base de datos
type Categorias struct {
	Id     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Status int64  `json:"status"`
}
