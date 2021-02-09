package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetTratamientosAnimal -v
func TestGetTratamientosAnimal(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/tratamientos", nil)
	req = mux.SetURLVars(req, map[string]string{"id_animal": "1"})
	response := execRequest(req, routes.GetTratamientosAnimal)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPostTratamientosAnimal -v
func TestPostTratamientosAnimal(t *testing.T) {

	var json = []byte(`[{"id_animal": 1, "nombre": "Prueba go test x2", "fecha": "2020-10-21"},{"id_animal": 5, "nombre": "Baño antipulgas pulgosas x3", "fecha": "2021-01-06"}]`)
	req, _ := http.NewRequest(http.MethodPost, "api/tratamientos", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PostTratamientosAnimal)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestDeleteTratamientosAnimal -v
func TestDeleteTratamientosAnimal(t *testing.T) {

	req, _ := http.NewRequest(http.MethodDelete, "api/tratamientos", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "14", "id_animal": "1"})
	response := execRequest(req, routes.DeleteTratamientosAnimal)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
