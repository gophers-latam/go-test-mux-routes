package tests

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestGetEstatusTodos -v
func TestGetEstatusTodos(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/estatus", nil)
	req = mux.SetURLVars(req, map[string]string{"tabla": "animales"})
	response := execRequest(req, routes.GetEstatusTodos)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
