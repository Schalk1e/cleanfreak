package plan

import (
	"errors"

	"github.com/spf13/cobra"
)

var PlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Prompts the user to construct a plan to execute.",
	Long: `You might want to defer actions to later, or perhaps not take action
on every file that is found. In this case, it is best to construct a clean plan
with this command to execute later.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("\n\n Remember to specify on object to construct a plan for")
	},
}
