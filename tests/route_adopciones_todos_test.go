package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/middlewares"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetAdopcionesTodos -v
func TestGetAdopcionesTodos(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/adopciones?p=1", nil)
	response := execRequest(req, routes.GetAdopcionesTodos)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetAdopcionesBusqueda -v
func TestGetAdopcionesBusqueda(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/adopciones/busqueda?p=1&n=tenebroso&e=gato", nil)
	response := execRequest(req, routes.GetAdopcionesBusqueda)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetAdopcionesDetalle -v
func TestGetAdopcionesDetalle(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/adopciones", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "11", "id_usuario": "1"})
	response := execRequest(req, routes.GetAdopcionesDetalle)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetAdopcionesAdmin -v
func TestGetAdopcionesAdmin(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/adopciones/admin/?p=1", nil)
	req = mux.SetURLVars(req, map[string]string{"id_admin": "1"})
	response := execRequest(req, routes.GetAdopcionesAdmin)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPostAdopcion -v
func TestPostAdopcion(t *testing.T) {

	var json = []byte(`{"id_usuario": 1, "nombre": "Supercan", "foto": "", "id_especie": "perro", "edad": "2 años", "folio": "AC-100", "sexo": "macho", "refugio": "Metropolis", "url": "https://form.joftorm.co/123456"}`)
	req, _ := http.NewRequest(http.MethodPost, "api/adopciones", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PostAdopcion)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPutAdopcion -v
func TestPutAdopcion(t *testing.T) {

	var json = []byte(`{"id": 12, "id_usuario": 1, "nombre": "Supercan updated", "foto": "super.png"}`)
	req, _ := http.NewRequest(http.MethodPut, "api/adopciones", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PutAdopcion)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestDeleteAdopcion -v
func TestDeleteAdopcion(t *testing.T) {

	req, _ := http.NewRequest(http.MethodDelete, "api/adopciones", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "7"})
	req.Header.Set("Authorization", testJWT)
	response := execRequest(req, middlewares.ValidarJWT(routes.DeleteAdopcion))

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
