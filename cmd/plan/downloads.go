package plan

import (
	"fmt"

	"github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var DownloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Runs a cleanfreak plan on the downloads folder.",
	Long: `
This command will prompt the user to construct a plan for whichever files are
found in the User's downloads folder. It will either save the plan to be applied
later, or it can be applied directly after the build with the apply flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		d := core.Dir{}

		// Steps to do:
		// Run filepicker once with Files to Delete as the title.
		// Run filepicker again with Files to Move as the title.
		// Ensure files to move don't contain any selections from files to
		// delete. (Use exclusions.)

		filepicker_model := cmdutil.FileTreeSelect(d.GetDownloads(), "Mark files for deletion:")

		fmt.Println(filepicker_model.SelectedFiles)
	},
}
