package tests

import (
	"net/http"
	"testing"

	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetDirectorioPerfiles -v
func TestGetDirectorioPerfiles(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/directorio?p=1", nil)
	response := execRequest(req, routes.GetDirectorioPerfiles)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetDirectorioBusqueda -v
func TestGetDirectorioBusqueda(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/directorio/busqueda?p=1&n=&t=servicio", nil)
	response := execRequest(req, routes.GetDirectorioBusqueda)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
