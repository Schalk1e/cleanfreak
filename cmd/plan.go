package cmd

import (
	"fmt"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Prompts the user to construct a plan (in yaml?) to execute later.",
	Long: `You might want to defer actions to later, or perhaps not take action
on every file that is found. In this case, it is best to construct a clean plan
with this command to execute later.`,
	Run: func(cmd *cobra.Command, args []string) {
		d := core.Dir{}
		// File select on downloads directory.
		m := cmdutil.FileTreeSelect(d.GetDownloads())

		fmt.Println(m.SelectedFiles)
		// At this point, we probably ask the user whether they'd like to audit
		// their selection, or if they're happy.
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
