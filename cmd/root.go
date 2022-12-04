package xpln

import (
	"fmt"
	"os"
	"strings"

	xpln "github.com/5amCurfew/xpln/pkg"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "xpln <PATH_TO_FILE>",
	Version: version,
	Short:   "xpln - a CLI to explain code blocks",
	Long:    `xpln is a command line interface to inspect and explain code blocks using OpenAI's GPT-3 (refer to beta.openai.com). `,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var file = args[0]
		var start, end string
		lineRange, _ := cmd.Flags().GetString("lines")

		// Check for lines flag provided
		if lineRange != "" {
			start = strings.Split(lineRange, "-")[0]
			end = strings.Split(lineRange, "-")[1]
		}

		var block = xpln.ReadCodeBlock(file, start, end)
		fmt.Println(block)
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("lines", "l", "", "Declare the lines of the code block (start-end e.g. 5-10)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error using xpln: '%s'", err)
		os.Exit(1)
	}
}
