# go-openai-sdk

This is an SDK for OpenAI's ChatCompletion API. It offers tools for doing API requests to execute chat completions.

HOW TO USE:
1. Create openaisdk APIClient instance. You will have to pass your API key here
2. Create chat conversation that you want the model to complete
3. Call the CreateChatCompletion method with desired AI model and chat conversation (messages) as arguments
4. Parse the response to desired format

EXAMPLE:

	// Create an instance of the APIClient from your SDK
	client := openaisdk.APIClient{APIKey: yourAPIKey}

	// Define the messages for the chat
	messages := []openaisdk.Message{
		{
			Role:    "system",
			Content: "You are a product search assistant. Your job is to suggest the types of products customers are looking for. You only offer outdoor equipment.",
		},
		{
			Role:    "user",
			Content: "I am looking for a backpack",
		},
		{
			Role:    "system",
			Content: "Good backpack is essential for a successful hiking experience. Here are some types of backpacks you might consider:\n1. Small day pack: Small and lightweight backpack is good for shorter day-hikes. Good backpack would be lightweight, water resistant and comfortable.\n 2. Hiking backpack: For longer hikes it is best to use a larger backpack. Good hiking backpack is around 65-85 liters in volume. It should have ample pouches and comfortable waist belt.",
		},
		{
			Role:    "user",
			Content: "I am looking for a sleeping pad",
		},
	}

	// Call the CreateChatCompletion method from your SDK
	responseBody := client.CreateChatCompletion("gpt-3.5-turbo", messages)
