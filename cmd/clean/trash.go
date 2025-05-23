package clean

import (
	"log"
	"runtime"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var TrashCmd = &cobra.Command{
	Use:   "trash",
	Short: "Runs cleanfreak on the trash folder.",
	Long:  "This command will clean the user's trash folder.",

	Run: func(cmd *cobra.Command, args []string) {
		// Always relative to home! Think abt this a little more carefully.
		diagnose_text := "No files in the Trash folder."

		d := core.Dir{}
		d.OS = runtime.GOOS

		if !cmdutil.IsIn(d.OS, cmdutil.TrashSupported[:]) {
			cmdutil.PrintWarning("Cleaning Trash not yet supported on this OS! Please contribute!")
			return
		}

		trashdir := d.GetTrash()
		if !core.DirExists(trashdir) {
			cmdutil.PrintDirectoryNotFound(trashdir)
			return
		}
		if core.DirEmpty(d.GetTrash()) {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			cmdutil.PrintOrder()
		} else {
			cmdutil.PrintDiagnoseFail(diagnose_text)
			trash := d.GetTrash()
			CleanTrash(trash)
			cmdutil.PrintOrder()
		}
	},
}

func CleanTrash(target string) {
	files := core.List(target, false)
	action := cmdutil.TrashDeleteSurvey()
	c := core.Clean{}
	if action == "Y" {
		for i := 1; i < len(files); i++ {
			c.SourceFile = files[i]
			if err := c.FileDelete(); err != nil {
				log.Printf("File not successfully removed: %v", err)
			}
		}
		cmdutil.PrintDeleted()
	}
}
