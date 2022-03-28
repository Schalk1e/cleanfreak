package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var principlesCmd = &cobra.Command{
	Use:   "principles",
	Short: "Returns 7 guiding principles of Cleanfreak.",
	Long: `
The principles command prints a list of the 10 guiding principles of Cleanfreak.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			`
1. No files in the Downloads folder.
2. No icons/files on Desktop.
3. No empty files or folders.
4. No large and unnecessary cache files. 
5. No unnested files in Documents.
6. Correct file types in associated folders.
7. An empty trash bin.`)
	},
}

func init() {
	listCmd.AddCommand(principlesCmd)
}
