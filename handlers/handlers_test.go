package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStatusHandler: Verifica se a rota /status devolve 200 e o JSON esperado
func TestStatusHandler(t *testing.T) {
	// Cria a requisição
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	//  Grava a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)
	handler.ServeHTTP(rr, req)

	// Valida o Código HTTP
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code errado: recebido %v esperado %v",
			status, http.StatusOK)
	}

	// Valida o Corpo (JSON)
	expected := `{"status":"ok","message":"servidor ativo"}`
	if rr.Body.String() != expected {
		t.Errorf("Body inesperado: recebido %v esperado %v",
			rr.Body.String(), expected)
	}
}

// TestAskHandler_Empty: Verifica se a API rejeita perguntas vazias (Erro 400)
func TestAskHandler_Empty(t *testing.T) {
	// Cria um JSON com pergunta vazia
	payload := []byte(`{"question": ""}`)

	req, err := http.NewRequest("POST", "/api/prompt", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AskHandler)
	handler.ServeHTTP(rr, req)

	// Esperamos erro 400 (Bad Request)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Deveria retornar erro 400 para pergunta vazia, mas retornou: %v", status)
	}
}
