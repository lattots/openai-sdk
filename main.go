package openaisdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIClient struct {
	APIKey string
}

type choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type ChatResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

type data struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []data `json:"data"`
	Model  string `json:"model"`
	Usage  usage  `json:"usage"`
}

func (c *APIClient) CreateChatCompletion(model string, messages []Message) (ChatResponse, error) {
	// URL for the openai's endpoint is set here
	apiURL := "https://api.openai.com/v1/chat/completions"

	// This is the payload of the API request
	payload := struct {
		Model    string    `json:"model"`
		Messages []Message `json:"messages"`
	}{
		Model:    model,
		Messages: messages,
	}

	// Payload is formed to JSON
	// This is done to pass it to openai API
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ChatResponse{}, err
	}

	// Request object "req" is created
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error forming the API request: ", err)
		return ChatResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// HTTP client is created
	client := http.Client{}
	// HTTP request is done
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing API request: ", err)
		return ChatResponse{}, err
	}
	defer resp.Body.Close()

	// Response's content is read to a variable
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return ChatResponse{}, err
	}

	// This is the response that is returned
	var response ChatResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Error unmarshalling the response: ", err)
		return ChatResponse{}, err
	}

	return response, nil
}

func (c *APIClient) CreateVectorEmbedding(model string, text string) (EmbeddingResponse, error) {
	// URL for the openai's endpoint is set here
	apiURL := "https://api.openai.com/v1/embeddings"

	// This is the payload of the API request
	payload := struct {
		Input string `json:"input"`
		Model string `json:"model"`
	}{
		Input: text,
		Model: model,
	}

	// Payload is formed to JSON
	// This is done to pass it to openai API
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return EmbeddingResponse{}, err
	}

	// Request object "req" is created
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error forming the API request: ", err)
		return EmbeddingResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// HTTP client is created
	client := http.Client{}
	// HTTP request is done
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing API request: ", err)
		return EmbeddingResponse{}, err
	}
	defer resp.Body.Close()

	// Response's content is read to a variable
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return EmbeddingResponse{}, err
	}

	var response EmbeddingResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Error unmarshalling the response: ", err)
		return EmbeddingResponse{}, err
	}

	return response, nil
}
