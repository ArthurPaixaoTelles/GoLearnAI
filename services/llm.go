package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func QueryLLM(question string) (string, error) {
	url := "https://router.huggingface.co/v1/chat/completions"
	model := "openai/gpt-oss-20b"

	apiKey := os.Getenv("HF_API_KEY")
	if apiKey == "" {
		return "", errors.New("HF_API_KEY não encontrado. Defina no .env ou variável de ambiente")
	}

	// Payload compatível com API atual
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": question},
		},
		"stream": false,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("erro da API: %s", string(body))
	}

	// Estrutura da resposta da API
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		OutputText string `json:"output_text"` // alternativo se existir
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Priorize choices[0].message.content
	if len(response.Choices) > 0 && response.Choices[0].Message.Content != "" {
		return response.Choices[0].Message.Content, nil
	}
	if response.OutputText != "" {
		return response.OutputText, nil
	}

	return "", errors.New("resposta inesperada da API")
}
