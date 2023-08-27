# OpenAI SDK for Go

This is an SDK for OpenAI's ChatCompletion API. It offers tools for doing API requests to execute chat completions.

##Chat Completions:
1. Create openaisdk APIClient instance. You will have to pass your API key here
2. Create chat conversation that you want the model to complete
3. Call the CreateChatCompletion method with desired AI model and chat conversation (messages) as arguments
4. Parse the response to desired format

###Example:

	// Create an instance of the APIClient from your SDK
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

	// Call the CreateChatCompletion method from your SDK
	responseBody := client.CreateChatCompletion("gpt-3.5-turbo", messages)

##Vector Embeddings
1. Create openaisdk APIClient instance. You will have to pass your API key here
2. Call the CreateVectorEmbedding method with desired AI model and text to be embedded as arguments
3. Parse the response to desired format

###Example:

	// Create an instance of the APIClient from your SDK
	client := openaisdk.APIClient{APIKey: yourAPIKey}

	text := "This is a test string for embeddings API"

	// Call the CreateVectorEmbedding method from your SDK
	responseBody = client.CreateVectorEmbedding("text-embedding-ada-002", text)