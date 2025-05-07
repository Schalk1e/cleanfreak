package clean

import (
	"errors"

	"github.com/spf13/cobra"
)

var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Initialises clean procedure for the specified item.",
	Long:  `This command will call the cleaning procedure for the specified item.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("\n\nRemember to specify an object to clean")
	},
}
