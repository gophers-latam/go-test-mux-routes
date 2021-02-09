package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/middlewares"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetAnimalesListado -v
func TestGetAnimalesListado(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/animales/listado/?p=1", nil)
	req = mux.SetURLVars(req, map[string]string{"id_usuario": "1"})
	response := execRequest(req, routes.GetAnimalesListado)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetAnimalesDetalle -v
func TestGetAnimalesDetalle(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/animales", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "5", "id_usuario": "1"})
	response := execRequest(req, routes.GetAnimalesDetalle)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetAnimalesBusqueda -v
func TestGetAnimalesBusqueda(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/animales/busqueda/?p=1&n=tene&e=", nil)
	req = mux.SetURLVars(req, map[string]string{"id_usuario": "1"})
	response := execRequest(req, routes.GetAnimalesBusqueda)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPostAnimal -v
func TestPostAnimal(t *testing.T) {

	var json = []byte(`{"id_usuario": 1, "nombre": "Tulkas", "id_especie": "perro", "cantidad": 1, "nacimiento": "2020-09-08", "raza": "Mestizo", "sexo": "macho", "ubicacion": "En Valinor XD", "estatus": "extraviado", "descripcion": "Es muy fuerte XD", "accesorio": "no", "foto": "tulkas_xd.png", "url": "https://form.tulkas.co/001", "recompensa": "si", "fecha_extravio": "2020-10-05"}`)
	req, _ := http.NewRequest(http.MethodPost, "api/animales", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PostAnimal)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPutAnimal -v
func TestPutAnimal(t *testing.T) {

	var json = []byte(`{"id": 6, "id_usuario": 1, "nombre": "El Señor Tenebroso xD2", "id_especie": "gato", "cantidad": 1, "nacimiento": "2020-10-21", "raza": "Mestizo", "sexo": "hembra", "ubicacion": "En algun lugar del inframundo XD2", "estatus": "extraviado", "descripcion": "Es muy bravo, puede invocar cosas XD2", "accesorio": "no", "foto": "señor_tenebroso_xd2.png", "url": "https://form.nagini.co/0012", "recompensa": "si", "fecha_extravio": "2020-11-21"}`)
	req, _ := http.NewRequest(http.MethodPut, "api/animales", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PutAnimal)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestDeleteAnimal -v
func TestDeleteAnimal(t *testing.T) {

	req, _ := http.NewRequest(http.MethodDelete, "api/animales", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "7"})
	req.Header.Set("Authorization", testJWT)
	response := execRequest(req, middlewares.ValidarJWT(routes.DeleteAnimal))

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
