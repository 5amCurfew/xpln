package xpln

import (
	"fmt"
	"os"
	"strings"

	xpln "github.com/5amCurfew/xpln/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var translateCmd = &cobra.Command{
	Use:     "translate <PATH_TO_FILE>",
	Aliases: []string{"t"},
	Short:   "Translate code block to another language",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var file = args[0]
		var start, end string
		lineRange, _ := cmd.Flags().GetString("lines")
		translateTo, _ := cmd.Flags().GetString("to")
		outputWidth := int(pterm.GetTerminalWidth() / 3)

		// Check for lines flag provided
		if lineRange != "" {
			start = strings.Split(lineRange, "-")[0]
			end = strings.Split(lineRange, "-")[1]
		}

		if translateTo == "" {
			fmt.Println("Failed xpln'ing! Please provide the language to translate to (--translate, -t)")
			os.Exit(1)
		}

		var block, _ = xpln.CreateCodeBlock(file, start, end)
		var translated, _ = xpln.TranslateCodeBlock(block, translateTo)

		panel1 := pterm.DefaultBox.WithTitle(
			pterm.LightWhite("Code Block")).Sprint(pterm.LightWhite(block.FormatBlockOutput(outputWidth)))
		panel2 := pterm.DefaultBox.WithTitle(
			pterm.Cyan("Translated")).Sprint(pterm.Cyan(xpln.Wrap(translated, outputWidth)))

		panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
			{{Data: panel1}, {Data: panel2}},
		}).Srender()

		pterm.DefaultBox.WithTitle(pterm.LightWhite("xpln")).Println(panels)
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
}
