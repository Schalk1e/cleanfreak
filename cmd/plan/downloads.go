package plan

import (
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
		// When running filepicker with Files to Move, loop over configured cf
		// base directories for now (Can expand on config options later.) Ensure
		// files are sequentially excluded from the set.

		// TODO
		// Support file exclusions.
		// Read from config here and support add Files to Move.
		// Add output to yaml plan.
		// Add state checker.
		// Add plan apply.

		p := PlanFiles{
			dir: d.GetDownloads(),
			// Hardcode this or now...
			move_dirs: []string{"a", "b"},
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

	},
}
