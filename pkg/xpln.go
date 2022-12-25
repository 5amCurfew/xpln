package xpln

import (
	"context"
	"os"
	"strings"

	util "github.com/5amCurfew/xpln/util"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/mitchellh/go-wordwrap"
)

// ///////////////////////////////////////////
// Create Code Block from file
// ///////////////////////////////////////////
func CreateCodeBlock(f, s, e string) (util.CodeBlock, error) {
	_, err := os.Open(f)

	if err != nil {
		return util.CodeBlock{}, err
	}

	return util.CodeBlock{
		File:    f,
		Lang:    util.DetermineLang(f),
		Comment: util.DetermineComment(f),
		Block:   util.ReadFile(f, s, e),
	}, nil
}

// ///////////////////////////////////////////
// Parse Code Block to OpenAI API
// ///////////////////////////////////////////
func ExplainCodeBlock(cb util.CodeBlock) (string, error) {

	openAPIKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(openAPIKey)

	var prefix = cb.Comment + " " + cb.Lang + " application\n"
	var suffix = "\n" + cb.Comment + " Here is what the " + cb.Lang + " application is doing:\n" + cb.Comment + " 1."
	prompt := prefix + cb.Block + suffix

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:           []string{prompt},
		MaxTokens:        gpt3.IntPtr(256),
		Stop:             []string{"\n\n"},
		TopP:             gpt3.Float32Ptr(0.05),
		FrequencyPenalty: 0.8,
	})

	if err != nil {
		return "GPT-3 Completion Error", err
	}

	return strings.Replace(string(cb.Comment+" 1."+resp.Choices[0].Text), cb.Comment+" ", "", -1), nil
}

func Format(e string, w int, limit bool) string {

	var lines = strings.Split(string(e), "\n")
	const maxOutput = 20

	var formatted string
	if limit {
		for i := 0; i < len(lines) && i <= maxOutput; i++ {
			if i == maxOutput {
				formatted += "..."
				break
			} else {
				formatted += lines[i] + "\n"
			}
		}
	} else {
		formatted = e
	}

	return wordwrap.WrapString(strings.Replace(formatted, "\t", "    ", -1), uint(w))
}
