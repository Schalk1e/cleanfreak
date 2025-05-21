package plan

import (
	"fmt"
	"os"

	"github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"path/filepath"
)

var apply bool

var DownloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Runs a cleanfreak plan on the downloads folder.",
	Long: `
This command will prompt the user to construct a plan for whichever files are
found in the User's downloads folder. It will either save the plan to be applied
later, or it can be applied directly after the build with the apply flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		var move_dirs []string

		d := core.Dir{}
		subdirs := viper.GetStringSlice("subdirs")
		rootdir := viper.GetString("directory")
		homedir, _ := os.UserHomeDir()

		for _, subdir := range subdirs {
			move_dirs = append(move_dirs, filepath.Join(homedir, rootdir, subdir))
		}

		p := PlanFiles{
			dir:       d.GetDownloads(),
			move_dirs: move_dirs,
		}

		p.ToDelete()
		p.ToMove()

		p.PrintPlan()

		// Ask here whether the user would like to save the plan or execute it
		// now.

		// If saved, we want a list of plans with a state that the user can
		// choose to view or execute. The lipgloss components would be amazing
		// for browsing and or editing plans.

		// Perhaps, for this PR we only add the option to execute immediately
		// and then we can deal with plan caching later.

		if apply {
			// Ask whether they want to apply
			choice := cmdutil.ListResult(
				[]string{"Y", "N"}, "Would you like to apply the plan now?",
			).Choice
			switch choice {
			case "Y":
				// Do deletes
				for _, file := range p.FilesToDelete {
					c := core.Clean{
						SourceFile: file,
					}
					err := c.FileDelete()
					if err != nil {
						fmt.Println("Error removing file: ", err)
					}
				}
				// Do moves
				for k, v := range p.FilesToMove {
					for _, file := range v {
						c := core.Clean{
							SourceFile: file,
							TargetFile: filepath.Join(k, filepath.Base(file)),
						}
						err := c.FileTransfer()
						if err != nil {
							fmt.Println("Error moving file: ", err)
						}
					}
				}
				// Print the executed plan here!
			case "N":
				fmt.Println("\nSkipping apply.")
			}
		}
	},
}

func init() {
	DownloadsCmd.Flags().BoolVar(
		&apply, "apply", false, "Whether to prompt the user to apply the plan.",
	)
}
