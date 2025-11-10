package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes define todas as rotas do sistema.
func RegisterRoutes(r *mux.Router) {
	// rota raiz
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// rota de status (pra teste)
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	r.HandleFunc("/api/prompt", AskHandler).Methods("POST")
}

// HomeHandler responde à rota raiz
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "🏠 Bem-vindo ao servidor Go! Essa tela está vindo do routes.go que está na pasta handlers")
}

// StatusHandler responde com um JSON de status simples
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"servidor ativo"}`))
}
