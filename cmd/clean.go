package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Initialises clean procedure for the specified item. (-A for all items).",
	Long:  `This command will call the cleaning procedure for the specified item.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
