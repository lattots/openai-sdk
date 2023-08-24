package openaisdk

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIClient struct {
	APIKey string
}

func (c *APIclient) CreateChatCompletion(model string, messages []Message) {
	// URL for the openai's endpoint is set here
	apiURL := "https://api.openai.com/v1/chat/completions"

	// This is the payload of the API request
	payload := struct {
		Model		string		'json:"model"'
		Messages	[]Message	'json:"messages"'
	}{
		Model:    model,
		Messages: messages,
	}

	// Payload is formed to JSON
	// This is done to pass it to openai API
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Request object "req" is created
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error forming the API request: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", '"Bearer"+c.APIkey')

	// HTTP client is created
	client := http.Client{}
	// HTTP request is done
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing API request: ", err)
		return
	}
	defer resp.Body.Close

	// Response's content is read to a variable
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	return responseBody
}