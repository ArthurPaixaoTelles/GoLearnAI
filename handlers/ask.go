package handlers

import (
	"encoding/json"
	"net/http"
	"project/services"
)

type QuestionRequest struct {
	Question string `json:"question"`
}

type AnswerResponse struct {
	Answer string `json:"answer"`
}

func AskHandler(w http.ResponseWriter, r *http.Request) {
	var req QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Question == "" {
		http.Error(w, "JSON inválido ou campo 'question' vazio", http.StatusBadRequest)
		return
	}

	// Chama o serviço que consulta o LLM
	answer, err := services.QueryLLM(req.Question)
	if err != nil {
		http.Error(w, "Erro ao consultar LLM: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a resposta
	json.NewEncoder(w).Encode(AnswerResponse{Answer: answer})
}
