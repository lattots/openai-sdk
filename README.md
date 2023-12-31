# OpenAI SDK for Go

This is an SDK for OpenAI's Chat and Embeddings APIs. It can be used to create chat completions and text embeddings with the OpenAI APIs.

## Chat Completions

1. Create openaisdk APIClient instance. You will have to pass your API key here
2. Create chat conversation that you want the model to complete
3. Call the CreateChatCompletion method with desired AI model and chat conversation (messages) as arguments
4. Access the values in response object

~~~go
// Create an instance of the APIClient
client := openaisdk.APIClient{APIKey: yourAPIKey}

// Define the messages for the chat
messages := []openaisdk.Message{
    {
        Role:    "system",
        Content: "You are a helpful assistant.",
    },
    {
        Role:    "user",
        Content: "Hello!",
    },
}

// Call the CreateChatCompletion method with your client
response, err := client.CreateChatCompletion("gpt-3.5-turbo", messages)
if err != nil {
    // Handle error
}

// You can access a single chat completion like this
completion := response.Choices[0].Message.Content
~~~

## Vector Embeddings

1. Create openaisdk APIClient instance. You will have to pass your API key here
2. Call the CreateVectorEmbedding method with desired AI model and text to be embedded as arguments
3. Access the values in response object

~~~go
// Create an instance of the APIClient
client := openaisdk.APIClient{APIKey: yourAPIKey}

text := "This is a test string for embeddings API"

// Call the CreateVectorEmbedding method with your client
response, err := client.CreateVectorEmbedding("text-embedding-ada-002", text)
if err != nil {
    // Handle error
}

// You can access a single embedding like this
vector := response.Data[0].Embedding
~~~
