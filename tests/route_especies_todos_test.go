package tests

import (
	"net/http"
	"testing"

	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetEspeciesTodos -v
func TestGetEspeciesTodos(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/especies", nil)
	response := execRequest(req, routes.GetEspeciesTodos)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
