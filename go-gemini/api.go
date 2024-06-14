package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type vulData struct {
	library_name      string
	vulnerability     string
	brief_description string
}

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	prompt := generatePrompt(vulData{
		library_name:  "@strapi/plugin-upload",
		vulnerability: "Denial-of-Service (DoS)",
		brief_description: `A Denial-of-Service was found in the media upload process causing the server to crash without restarting,
							affecting either development and production environments.DetailsUsually, 
							errors in the application cause it to log the error and keep it running for other clients. 
							This behavior, in contrast, stops the server execution, making it unavailable for any clients until it's manually restarted.`,
	})

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range resp.Candidates {
		fmt.Println(v.Content)
	}
}

func generatePrompt(data vulData) string {
	return fmt.Sprintf("Summarize the following text: It's important to patch these vulnerabilities to protect systems from attacks.\n"+
		"The format of the first sentence must be '%s is vulnerable to %s.'\n"+
		"The second sentence must be 'The vulnerability is due to %s due to %s.'",
		data.library_name, data.vulnerability, data.vulnerability, data.brief_description)
}
