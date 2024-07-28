package openaisdk

import (
	"os"
	"testing"
)

func TestCreateChatCompletion(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	client := NewAPIClient(apiKey)

	messages := []Message{
		{
			Role: "user",
			Content: []Content{
				NewTextContent("What's in this image?"),
				NewImageContent("https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"),
			},
		},
	}

	response, err := client.CreateChatCompletion("gpt-4o", messages, 300)
	if err != nil {
		t.Fatalf("CreateChatCompletion error: %v", err)
	}
	if response.Id == "" {
		t.Errorf("Expected non-empty response ID")
	}
	if len(response.Choices) == 0 {
		t.Errorf("Expected at least one choice in response")
	}
	if response.Choices[0].Message.Content == "" {
		t.Errorf("Expected non-empty message content")
	}
}

func TestCreateVectorEmbedding(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	client := NewAPIClient(apiKey)

	response, err := client.CreateVectorEmbedding("text-embedding-ada-002", "OpenAI is amazing.")
	if err != nil {
		t.Fatalf("CreateVectorEmbedding error: %v", err)
	}

	if response.Object != "list" {
		t.Errorf("Expected response object to be 'list', got %s", response.Object)
	}
	if len(response.Data) == 0 {
		t.Errorf("Expected non-empty data in response")
	}
	if len(response.Data[0].Embedding) == 0 {
		t.Errorf("Expected non-empty embedding in response data")
	}
}
