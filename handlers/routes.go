package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	//  Rotas da API (Têm prioridade)
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	r.HandleFunc("/api/prompt", AskHandler).Methods("POST")

	//  Rota para servir o Front-end (Arquivos Estáticos)
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)
}

// StatusHandler responde com um JSON de status simples
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"servidor ativo"}`))
}
