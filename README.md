# CRUD con Golang

## Este es un proyecto de operaciones basicas sobre la tabla Products, para ejecutar el programa necesitamos seguir los siguientes pasos:

- Instalar MariaDB(https://mariadb.org/download/)
- Clonar este repositorio
  - ```git clone https://github.com/smithasencios/CRUD_With_Golang.git```
- Instalar git, dep(https://github.com/golang/dep)
- Cuando Instalar MariaDB, tambien te instala un IDE llamado HeidiSQL, ahi ejecutar los archivos que estan en la carpeta database:
  - ```database_data_script```
  - ```database_creation_script.sql```
- Abrir el proyecto con Visual Studio Code, abrir el terminal y ejecutar:
  - ```dep ensure```
- Para compilar el proyecto ejecutar:
  - ```go build .```
- Buscamos el archivo generado, tiene el nombre del proyecto con la extension .exe, ejecutamos este archivo, y podemos empezar a consumir los recursos.

## Recursos
### GET (Obtener todos los productos)
- url : http://localhost:8005/products
### POST (Agregar Producto)
- url : http://localhost:8005/products
- body: { "product_code":"XXX1","description":"DEscription 1" }
### PUT (Actualizar Producto)
- url : http://localhost:8005/products/100
- body: { "product_code":"XXX1","description":"DEscription actualizada" }
### Delete (Eliminar Producto)
- url : http://localhost:8005/products/100
