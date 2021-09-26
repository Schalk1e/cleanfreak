package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var principlesCmd = &cobra.Command{
	Use:   "principles",
	Short: "Returns 10 guiding principles of Cleanfreak.",
	Long: `
The principles command prints a list of the 10 guiding principles of Cleanfreak.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			`
1. No files in the Downloads folder.
2. No icons on Desktop.
3. No empty files or folders.
4. No large and unnecessary cache files. 
5. No unnested files in Documents.
6. Consistent folder naming convention.
7. Correct file types in associated folders.
8. No unneeded install files.
9. An empty trash bin.
10. Preferably, no unused apps.`)
	},
}

func init() {
	listCmd.AddCommand(principlesCmd)
}
