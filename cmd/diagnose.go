package cmd

import (
	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check compliance of workspace.",
	Long: `
The diagnose command will check whether the host workspace complies with each of the 10 guiding
principles of a clean workspace without taking any explicit action. The results will be reported
back to the user.`,

	Run: func(cmd *cobra.Command, args []string) {
		if core.DirEmpty("Downloads") {
			cmdutil.PrintDiagnoseSuccess("No files in the Downloads folder.")
		} else {
			cmdutil.PrintDiagnoseFail("No files in the Downloads folder.")
		}

		if core.DirEmpty("Desktop") {
			cmdutil.PrintDiagnoseSuccess("No icons/files on Desktop.")
		} else {
			cmdutil.PrintDiagnoseFail("No icons/files on Desktop.")
		}
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
