package xpln

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "c",
	Aliases: []string{"r"},
	Short:   "Do something else",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Doing something else")
	},
}

func init() {
	rootCmd.AddCommand(cmd)
}
