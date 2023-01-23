package cmd

import (
	"fmt"

	"github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Runs cleanfreak on the user's cache folder.",
	Long:  `This command will clean the user's cache folder.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nCache called")
	},
}

func PrintCache(dir string) {
	core.List(dir, false)

}

func init() {
	cleanCmd.AddCommand(cacheCmd)
}
