package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Runs cleanfreak on the downloads folder.",
	Long: `This command will clean the user's downloads folder by prompting the user
	to either transfer those files to the appropriate location in the cleanfreak project directory,
	or remove them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Downloads called")
	},
}

func init() {
	cleanCmd.AddCommand(downloadsCmd)
}
