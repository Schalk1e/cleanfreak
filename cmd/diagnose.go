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
		if downDirEmpty() {
			cmdutil.PrintDiagnoseSuccess()
		} else {
			cmdutil.PrintDiagnoseFail()
		}
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}

func downDirEmpty() (is_empty bool) {
	var dir string
	var dir_type = "downloads"

	dir = core.Dir(dir_type)
	is_empty = core.IsEmpty(dir)

	return
}
