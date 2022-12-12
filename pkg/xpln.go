package xpln

import (
	"context"
	"os"
	"strings"

	util "github.com/5amCurfew/xpln/util"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

// ///////////////////////////////////////////
// Read Code Block from file
// ///////////////////////////////////////////
func CreateCodeBlock(file, start, end string) (util.CodeBlock, error) {
	return util.NewCodeBlock(file, start, end)
}

// ///////////////////////////////////////////
// Parse Code Block to OpenAI API
// ///////////////////////////////////////////
func ExplainCodeBlock(block util.CodeBlock) (string, error) {
	godotenv.Load()

	openAPIKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(openAPIKey)

	var prefix = block.GetComment() + " " + block.GetLang() + " application\n"
	var suffix = "\n" + block.GetComment() + " Explain what the the application is doing:\n" + block.GetComment() + " 1."
	prompt := prefix + block.GetCode() + suffix

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:           []string{prompt},
		MaxTokens:        gpt3.IntPtr(128),
		Stop:             []string{"\n\n"},
		TopP:             gpt3.Float32Ptr(0.05),
		FrequencyPenalty: 0.8,
	})

	if err != nil {
		return "GPT-3 Completion Error", err
	}

	return strings.Replace(string(block.GetComment()+" 1."+resp.Choices[0].Text), block.GetComment()+" ", "", -1), nil
}
