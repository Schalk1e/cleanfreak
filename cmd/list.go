package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of an object specified by the user.",
	Long: `
The list command simply returns a list of items in an object specified by the user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("\n\nRemember to specify an object to list")
	},
}
