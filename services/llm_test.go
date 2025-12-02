package services

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestQueryLLM_SimulatedSuccess(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validar Header
		if r.Header.Get("Authorization") == "" {
			t.Errorf("Esperava header Authorization, mas veio vazio")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"choices": [
				{
					"message": {
						"content": "Eu sou uma IA simulada do GPT!"
					}
				}
			]
		}`))
	}))
	defer mockServer.Close()

	// Troca de URL (Injection)
	oldURL := HuggingFaceURL
	HuggingFaceURL = mockServer.URL
	defer func() { HuggingFaceURL = oldURL }()

	// Configura chave fake
	os.Setenv("HF_API_KEY", "chave_teste")
	defer os.Unsetenv("HF_API_KEY")

	// Executa
	resposta, err := QueryLLM("Teste")

	// Valida
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	esperado := "Eu sou uma IA simulada do GPT!"
	if resposta != esperado {
		t.Errorf("Resposta errada. Esperava '%s', recebeu '%s'", esperado, resposta)
	}
}
