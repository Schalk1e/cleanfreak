package cmd

import (
	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Returns the different items that cleanfreak will check for compliance.",
	Long: `
The items command will list the different items that cleanfreak will check for compliance.`,
	Run: func(cmd *cobra.Command, args []string) {

		plst := []string{"No files in the Downloads folder.", "No icons/files on Desktop.", "No large and unnecessary cache files.", "An empty trash bin."}

		cmdutil.PrintArrows(plst)

	},
}

func init() {
	listCmd.AddCommand(itemsCmd)
}
