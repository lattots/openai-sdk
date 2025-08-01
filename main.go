package openaisdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ContentType represents the type of content in the message
type ContentType string

const (
	TextType  ContentType = "text"
	ImageType ContentType = "image_url"
)

// TextContent represents text content in the message
type TextContent struct {
	Type ContentType `json:"type"`
	Text string      `json:"text,omitempty"`
}

// ImageContent represents image URL content in the message
type ImageContent struct {
	Type     ContentType `json:"type"`
	ImageUrl ImageUrl    `json:"image_url,omitempty"`
}

type ImageUrl struct {
	Url    string `json:"url"`
	Detail string `json:"detail,omitempty"`
}

// Message represents a message in the conversation
type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

// Content represents either text or image content
type Content any

func NewTextContent(text string) TextContent {
	return TextContent{
		Type: TextType,
		Text: text,
	}
}

func NewImageContent(url string, detail ...string) ImageContent {
	// Detail parameter is provided, it is attached to request
	var d string
	if len(detail) == 1 {
		d = detail[0]
	}
	return ImageContent{
		Type:     ImageType,
		ImageUrl: ImageUrl{Url: url, Detail: d}, // If detail is empty, it is omitted from json
	}
}

// APIClient represents a client for OpenAI API
type APIClient struct {
	APIKey string
	APIURL string
}

func NewAPIClient(apiKey string) *APIClient {
	return &APIClient{
		APIKey: apiKey,
		APIURL: "https://api.openai.com/v1",
	}
}

type responseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type choice struct {
	Index        int             `json:"index"`
	Message      responseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
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

func (c *APIClient) CreateChatCompletion(model string, messages []Message, maxTokens int) (ChatResponse, error) {
	apiURL := c.APIURL + "/chat/completions"

	payload := struct {
		Model     string    `json:"model"`
		Messages  []Message `json:"messages"`
		MaxTokens int       `json:"max_tokens"`
	}{
		Model:     model,
		Messages:  messages,
		MaxTokens: maxTokens,
	}

	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(payload)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error encoding JSON: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, &requestBody)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error forming the API request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error executing API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ChatResponse{}, fmt.Errorf("error executing API request: %s", resp.Status)
	}

	var response ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("error decoding response body: %w", err)
	}

	return response, nil
}

func (c *APIClient) CreateVectorEmbedding(model string, text string) (EmbeddingResponse, error) {
	apiURL := c.APIURL + "/embeddings"

	payload := struct {
		Input string `json:"input"`
		Model string `json:"model"`
	}{
		Input: text,
		Model: model,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return EmbeddingResponse{}, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return EmbeddingResponse{}, fmt.Errorf("error forming the API request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return EmbeddingResponse{}, fmt.Errorf("error executing API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return EmbeddingResponse{}, fmt.Errorf("error executing API request: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EmbeddingResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	var response EmbeddingResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return EmbeddingResponse{}, fmt.Errorf("error unmarshalling the response: %w", err)
	}

	return response, nil
}
