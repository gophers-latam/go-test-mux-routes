
# Ejemplos *go test* http mux routes 

Teniendo una estructura de proyecto similar de REST api, o no. Los test se ejecutan independientemente de si el proyecto completo este en ejecución.

![captura](https://awebytes.files.wordpress.com/2021/02/captura.png)

En los test se cubren ejemplo varios sobre GET, POST, PUT, DELETE y pasar JWT


- file_name **_test** .go => nomenclatura requerida de los archivos para go test


### Comandos esenciales:

**go test** => para correr todas las pruebas definidas

**go test -v** => para detalles de ejecución de las pruebas

**go test -run TestNameFunc** => para ejecutar una prueba en especifico de todas la definidas

**go test -run TestNameFunc -v** => para lo anterior con detalles