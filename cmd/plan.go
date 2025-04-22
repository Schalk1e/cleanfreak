package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Prompts the user to construct a plan (in yaml?) to execute later.",
	Long: `You might want to defer actions to later, or perhaps not take action
on every file that is found. In this case, it is best to construct a clean plan
with this command to execute later.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Plan called.")
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
