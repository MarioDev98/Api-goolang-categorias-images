// METODOS PARA LAS CONSULTAS DE BASE DE DATOS
package main

// todods los metodos de la api stories
// METODO POST
func createImages(stories Stories) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO imagenes (nombre,categoria,ruta) VALUES (?,?,?)", stories.Nombre, stories.Categoria, stories.Ruta)
	return err
}

// METODO PUT
// desactiva o activa el status
// 1 activo - 0 desactivo
func updateImages(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE imagenes SET  status = 0 WHERE id = ?", id)
	return err
}

// METODO GET
// trae la lista de todas las stories que existen
func getImages() ([]Stories, error) {
	//declare array por en caso de erro ser guarde el erroe
	stories := []Stories{}
	bd, err := getDB()
	if err != nil {
		return stories, err
	}
	// itera la base de datos
	rows, err := bd.Query("SELECT id, nombre, categoria, ruta,status FROM imagenes")
	if err != nil {
		return stories, err
	}

	for rows.Next() {

		var storie Stories
		err = rows.Scan(&storie.Id, &storie.Nombre, &storie.Categoria, &storie.Ruta, &storie.Status)
		if err != nil {
			return stories, err
		}
		// se llena el areglo
		stories = append(stories, storie)
	}
	return stories, nil
}

// METODO GET+ID
// este metodo solo trae la storie que consultemos aunque segun yo no la ocupo porque ya tengo el buscador
func getImageById(id int64) (Stories, error) {
	var categorias Stories
	bd, err := getDB()
	if err != nil {
		return categorias, err
	}
	row := bd.QueryRow("SELECT id, nombre,categoria, ruta, status FROM imagenes WHERE id = ?", id)
	err = row.Scan(&categorias.Id, &categorias.Nombre, &categorias.Categoria, &categorias.Ruta, &categorias.Status)
	if err != nil {
		return categorias, err
	}
	// Success!
	return categorias, nil
}
