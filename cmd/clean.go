package cmd

import (
	"fmt"

	"errors"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Initialises clean procedure for the specified item.",
	Long:  `This command will call the cleaning procedure for the specified item.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Clean called.")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("\n\nRemember to specify an object to clean!\n")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
