package main

// METODO DE API CATEGORIAS
// MEOTO GET DE CATEGORIAS
func createCategorias(categoria Categorias) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO categorias (nombre) VALUES (?)", categoria.Nombre)
	return err
}

func deleteCategorias(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM categorias WHERE id = ?", id)
	return err
}

// METODO PUT
func updateCategorias(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE categorias SET status = 0 WHERE id = ?", id)
	return err
}
func activar(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE categorias SET status = 1 WHERE id = ?", id)
	return err
}

// func updateCategorias(categoria Categorias) error {
// 	bd, err := getDB()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = bd.Exec("UPDATE categorias SET nombre = ?, status = ? WHERE id = ?", categoria.Nombre, categoria.Status, categoria.Id)
// 	return err
// }

// METODO GET
func getCategorias() ([]Categorias, error) {
	//declare array por en caso de erro ser guarde el erroe
	categorias := []Categorias{}
	bd, err := getDB()
	if err != nil {
		return categorias, err
	}
	// itera la base de datos
	rows, err := bd.Query("SELECT id, nombre, status FROM categorias ORDER BY id desc")
	if err != nil {
		return categorias, err
	}

	for rows.Next() {

		var categoria Categorias
		err = rows.Scan(&categoria.Id, &categoria.Nombre, &categoria.Status)
		if err != nil {
			return categorias, err
		}
		// se llena el areglo
		categorias = append(categorias, categoria)
	}
	return categorias, nil
}

// METODO GET + ID
// aunque tambien siento que no lo ocupo porque ya tengo el buscador pero todopuede pasar eh
func getCategoriaById(id int64) (Categorias, error) {
	var categorias Categorias
	bd, err := getDB()
	if err != nil {
		return categorias, err
	}
	row := bd.QueryRow("SELECT id, nombre, status FROM categorias WHERE id = ?", id)
	err = row.Scan(&categorias.Id, &categorias.Nombre, &categorias.Status)
	if err != nil {
		return categorias, err
	}
	// Success!
	return categorias, nil
}
