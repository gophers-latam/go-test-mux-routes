package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/backend/routes"
)

// cmd: go test -run TestPostUsuarioRegistro -v
func TestPostUsuarioRegistro(t *testing.T) {

	var json = []byte(`{"admin": 0, "username": "test", "nombre": "test", "email": "test@mail.com", "password": "123456"}`)
	req, _ := http.NewRequest(http.MethodPost, "api/usuarios/nuevo", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PostUsuarioRegistro)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPostUsuarioLogin -v
func TestPostUsuarioLogin(t *testing.T) {

	var json = []byte(`{"email": "zero@mail.com", "password": "123456"}`)
	req, _ := http.NewRequest(http.MethodPost, "api/usuarios/login", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PostUsuarioLogin)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestGetUsuarioPerfil -v
func TestGetUsuarioPerfil(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "api/usuarios", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	response := execRequest(req, routes.GetUsuarioPerfil)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Logf("Respuesta: %s", body)
	}
}

// cmd: go test -run TestPutUsuario -v
func TestPutUsuario(t *testing.T) {

	var json = []byte(`{"id": 3, "email": "test@mail.com", "tipo": "servicio", "foto": "test.png", "biografia": "test", "ubicacion": "test lugar", "redes": "fb.com/test", "numeros": "999123469"}`)
	req, _ := http.NewRequest(http.MethodPut, "api/usuarios", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	response := execRequest(req, routes.PutUsuario)

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Logf("Respuesta: %s", body)
	}
}
