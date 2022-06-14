package cmd

import (
	"fmt"
	"runtime"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check compliance of workspace.",
	Long: `
The diagnose command will check whether the host workspace complies with each of the items
shown by cf list items without taking any explicit action. The results will be reported
back to the user.`,

	Run: func(cmd *cobra.Command, args []string) {
		d := core.Dir{}
		d.OS = runtime.GOOS

		dirs := []string{d.GetDownloads(), d.GetDesktop(), d.GetTrash()}
		dirnames := []string{"Downloads", "Desktop", "Trash"}

		trash_supported := cmdutil.IsIn(d.OS, cmdutil.TrashSupported[:])

		if !trash_supported {
			dirs = dirs[:2]
			dirnames = dirnames[:2]
		}
		for indx, dir := range dirs {
			if core.DirEmpty((dir)) {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseSuccess(text)
			} else {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseFail(text)
			}
		}
		if !trash_supported {
			cmdutil.PrintWarning("Cleaning Trash not yet supported on this OS! Please contribute!")
		}
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
