package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}

var testJWT string = "Bearereykey..."

// util test funcs routes
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Código de respuesta esperado: %d - Obtenido: %d\n", expected, actual)
	}
}

func execRequest(r *http.Request, f func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(f)
	handler.ServeHTTP(rec, r)
	return rec
}
