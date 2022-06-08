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

		dirs := [3]string{d.GetDownloads(), d.GetDesktop(), d.GetTrash()}
		dirnames := [3]string{"Downloads", "Desktop", "Trash"}

		for indx, dir := range dirs {
			if core.DirEmpty((dir)) {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseSuccess(text)
			} else {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseFail(text)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
