package xpln

import (
	"context"
	"log"
	"os"

	util "github.com/5amCurfew/xpln/util"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

// ///////////////////////////////////////////
// Read Code Block from file
// ///////////////////////////////////////////
func CreateCodeBlock(file, start, end string) util.CodeBlock {
	return util.NewCodeBlock(file, start, end)
}

// ///////////////////////////////////////////
// Parse Code Block to OpenAI API
// ///////////////////////////////////////////
func ExplainCodeBlock(block util.CodeBlock) string {
	godotenv.Load()

	openAPIKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(openAPIKey)

	var prefix = block.GetComment() + " " + block.GetLang() + "\n" + block.GetComment() + " Code Block\n"
	var suffix = "\n" + block.GetComment() + " Here's what the Code-Block is doing:\n" + block.GetComment() + " 1."
	prompt := prefix + block.GetCode() + suffix

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:      []string{prompt},
		MaxTokens:   gpt3.IntPtr(256),
		Stop:        []string{"\n\n"},
		TopP:        gpt3.Float32Ptr(1),
		Temperature: gpt3.Float32Ptr(0.25),
	})

	if err != nil {
		log.Fatalln(err)
	}

	return string(block.GetComment() + " 1." + resp.Choices[0].Text)
}
