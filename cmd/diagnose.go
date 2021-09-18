package cmd

import (
	"fmt"

	core "github.com/Schalk1e/cleanfreak/core"
	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
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
		var cyan = cmdutil.Cyan()
		var green = cmdutil.Green()
		var red = cmdutil.Red()
		var bold = cmdutil.Bold()
		var end = cmdutil.End()

		if downDirEmpty() {
			fmt.Println(bold + cyan + "Done: " + end + "No files in the Downloads folder." + green + "✔" + end)
		} else {
			fmt.Println(bold + cyan + "Done: " + end + "No files in the Downloads folder." + red + "✘" + end)
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
