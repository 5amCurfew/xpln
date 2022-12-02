package xpln

import (
	"fmt"
	"os"

	xpln "github.com/5amCurfew/xpln/pkg"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "xpln",
	Version: version,
	Short:   "xpln - a CLI to explain code blocks",
	Long:    `xpln is a command line interface to inspect and explain code blocks using OpenAI's GPT-3 (refer to beta.openai.com). `,
	Args:    cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		var file = args[0]
		var start, end string
		if len(args) == 2 {
			start = args[1]
		} else if len(args) == 3 {
			start = args[1]
			end = args[2]
		}

		var block = xpln.ReadCodeBlock(file, start, end)
		fmt.Println(block)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error using xpln: '%s'", err)
		os.Exit(1)
	}
}
