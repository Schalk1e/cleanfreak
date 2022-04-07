package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs cleanfreak across all items.",
	Long: `This command will execute the cleanfreak process over all items 
	listed by cleanfreak list items.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("All called")
	},
}

func init() {
	cleanCmd.AddCommand(allCmd)
}
