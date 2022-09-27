# apicategorias
 Creación de una API REST con Golang/Go y MySQL

## Instalacion
* Clonar o descargar
* Intalar las dependencias (copiar y pegar en terminal dentro de la ubicacion del proyecto)
```
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
go get github.com/gorilla/mux

```
* Crear la Base de datos en MYSQL con el nombre TMAK
* modificar o crear el archivo  **.env** basado al del poryecto **.env**
* En el archivo  **.env** van las credenciales de conexion
* Creada la base de datos, exportar el archivo que se presenta de la tabla categorias.sql
* Compilar el proyecto para obtener el ejecutable con el comando  `go build`
* Ejecutar el archivo con  **apicategorias.exe** 
si pide permisos darselos

* Probar en postman o consumirla en su proyecto, estara activo en el la ruta
localhost:8000 + el nombre de la ruta en el archivo routes.go
 
 ## RUTAS

 # GET
 * localhost:8000/stories/
 
 # GET ID
  * localhost:8000/stories/1
 
 # POST
  * localhost:8000/stories/
 
 # PUT
 * localhost:8000/stories/1


