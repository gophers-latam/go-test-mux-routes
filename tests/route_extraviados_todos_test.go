package tests

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetExtraviadosTodos -v
func TestGetExtraviadosTodos(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/extraviados?p=1", nil)
	response := execRequest(req, routes.GetExtraviadosTodos)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetExtraviadosAnimal -v
func TestGetExtraviadosAnimal(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/extraviados", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	response := execRequest(req, routes.GetExtraviadosAnimal)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetExtraviadosBusqueda -v
func TestGetExtraviadosBusqueda(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/extraviados/busqueda?p=1&n=&m=11&e=gato", nil)
	response := execRequest(req, routes.GetExtraviadosBusqueda)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
