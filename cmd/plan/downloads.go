package plan

import (
	"fmt"

	"github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		d := core.Dir{}
		subdirs := viper.GetStringSlice("subdirs")

		p := PlanFiles{
			dir:       d.GetDownloads(),
			move_dirs: subdirs,
		}

		p.ToDelete()
		p.ToMove()

		// Print plan for the user
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
				// Do plan
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
