package cmd

import (
	"fmt"
	"runtime"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var trashCmd = &cobra.Command{
	Use:   "trash",
	Short: "Runs cleanfreak on the downloads folder.",
	Long:  "This command will clean the user's trash folder.",

	Run: func(cmd *cobra.Command, args []string) {
		// Always relative to home! Think abt this a little more carefully.
		// homedir, err := os.UserHomeDir() // Add error handling.
		diagnose_text := "No files in the Trash folder."

		d := core.Dir{}
		d.OS = runtime.GOOS

		if core.DirEmpty(d.GetTrash()) {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			fmt.Println("\nEverything is in order! 🎉")
			return
		}
	},
}

func init() {
	cleanCmd.AddCommand(trashCmd)
}
