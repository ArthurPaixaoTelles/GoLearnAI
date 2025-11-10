package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega o .env (opcional)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: .env não encontrado, usando valores padrão.")
	}

	// Cria o roteador
	r := mux.NewRouter()

	// Registra as rotas a partir do pacote handlers
	handlers.RegisterRoutes(r)

	// Define porta (default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor rodando na porta %s 🚀\n", port)
	link := fmt.Sprintf("http://localhost:%s", port)
	fmt.Printf("Servidor rodando! Acesse: %s 🚀\n", link)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
