package xpln

import (
	"fmt"
	"os"
	"strings"

	xpln "github.com/5amCurfew/xpln/pkg"
	"github.com/pterm/pterm"
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

		spinner, _ := pterm.DefaultSpinner.Start("xpln'ing " + file)
		spinner.SuccessPrinter = &pterm.PrefixPrinter{
			MessageStyle: &pterm.Style{pterm.BgBlack, pterm.FgCyan},
			Prefix: pterm.Prefix{
				Style: &pterm.Style{pterm.BgBlack, pterm.FgCyan},
				Text:  "✓",
			},
		}
		spinner.FailPrinter = &pterm.PrefixPrinter{
			MessageStyle: &pterm.Style{pterm.BgBlack, pterm.FgCyan},
			Prefix: pterm.Prefix{
				Style: &pterm.Style{pterm.BgBlack, pterm.FgLightRed},
				Text:  "x",
			},
		}

		var block, err = xpln.CreateCodeBlock(file, start, end)
		if err != nil {
			spinner.Fail("Failed xpln'ing! (" + err.Error() + ")")
			os.Exit(1)
		} else {
			var explained, err = xpln.ExplainCodeBlock(block)

			if err != nil {
				spinner.Fail("Failed xpln'ing! (" + err.Error() + ")")
				os.Exit(1)
			}

			panel1 := pterm.DefaultBox.WithTitle(pterm.LightWhite("Code Block")).Sprint(pterm.LightWhite(block.FormatBlock()))
			panel2 := pterm.DefaultBox.WithTitle(pterm.Cyan("Explained")).Sprint(pterm.Cyan(explained))

			panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
				{{Data: panel1}, {Data: panel2}},
			}).Srender()

			spinner.Success("Xpln'd!")

			pterm.DefaultBox.WithTitle(pterm.LightWhite("xpln")).Println(panels)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("lines", "l", "", "Declare the lines of the code block (start-end e.g. 5-10)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error using xpln: '%s'", err)
		os.Exit(1)
	}
}
