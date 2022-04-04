package cmd

import (
	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/spf13/cobra"
)

var principlesCmd = &cobra.Command{
	Use:   "principles",
	Short: "Returns 7 guiding principles of Cleanfreak.",
	Long: `
The principles command prints a list of the 10 guiding principles of Cleanfreak.`,
	Run: func(cmd *cobra.Command, args []string) {

		plst := []string{"No files in the Downloads folder.", "No icons/files on Desktop.", "No empty files or folders.", "No large and unnecessary cache files.", "No unnested files in Documents.", "Correct file types in associated folders.", "An empty trash bin."}

		cmdutil.PrintArrows(plst)

	},
}

func init() {
	listCmd.AddCommand(principlesCmd)
}
