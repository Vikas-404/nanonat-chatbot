// the package below is a package under textToText in the folder open editor
// the main function of this package entirely is to make a POST request to openAI 's backend API,
// and get response from the openAI's backend API and gives the response to the user
package textToText

//the import below is where all the packages needed for this program is imported
import (
	//this packages below are local packages that came pre-installed with go language, and are used to run the program
	"context"
	"fmt"
	"strings"
	"os"
	"log"

	//this is the open ai that is use to make a request
	openai "github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
	

)

// the variable below is used to store previous messages so that it can rememeber previous chat
// the below is a variable chatHistory of type []openai.ChatCompletionMessage
var chatHistory []openai.ChatCompletionMessage

// the ToText function below sends a POST request to openAI using my private API_KEY and returns a string as response
// The ToText function takes in a string as an argument and returns a string as a response else if response was nil it will send an error message
func ToText(inputText string) string {
	// You can find your API key at https://platform.openai.com/account/api-keys

	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	client := openai.NewClient(apiKey)
	// saving previous chat to recent chat to remember the chat
	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: inputText,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: chatHistory,
		},
	)
	//this is just an error message that is being returned if the system is not connected to the internet cause its running on local server
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "Back-End server not connected to the internet. Please make sure that you're connected to the internet."
	}

	responseText := resp.Choices[0].Message.Content
	// where some text was modifies and filtered
	//here we modify the string to what you want that will give the user(pre-trained)
	modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Vikas")
	modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Silver Oak University Computer Science Students")
	modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Vikas")

	//the below is the recent chat being appended or added to the previous chat so that it can give you better and updated
	// response and also can remember the previous chat
	chatHistory = append(chatHistory, resp.Choices[0].Message)

	//this is the return of the string above
	return modifiedResponse2
}
