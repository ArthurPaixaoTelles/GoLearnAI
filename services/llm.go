package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func QueryLLM(question string) (string, error) {
	// Endpoint gratuito da HuggingFace
	url := "https://api-inference.huggingface.co/models/google/flan-t5-base"

	apiKey := os.Getenv("HF_API_KEY")
	if apiKey == "" {
		return "", errors.New("chave da API não encontrada no .env")
	}

	// Corpo da requisição
	body := map[string]string{"inputs": question}
	jsonData, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("erro da API: %s", string(body))
	}

	var response []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	// Extrai texto da resposta
	if len(response) > 0 {
		generatedText, ok := response[0]["generated_text"].(string)
		if ok {
			return generatedText, nil
		}
	}

	return "", errors.New("resposta inesperada da API")
}
