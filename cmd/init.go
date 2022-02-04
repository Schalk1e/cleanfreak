package cmd

import (
	"os"

	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a cleanfreak folder.",
	Long: `This command will initialise a default cleanfreak folder in the specified
	location that contains a number of subfolders that are intended to provide appropriate homes
	for most filetypes.`,
	Run: func(cmd *cobra.Command, args []string) {
		base_dir, _ := cmd.Flags().GetString("path")

		subdir := []string{"personal", "work"}
		// What happens if these already exist? MkdirAll will return nil. Be verbose about this...
		if base_dir == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			} else {
				core.DirAdd(homedir, subdir)
			}
		} else {
			core.DirAdd(base_dir, subdir)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("path", "", "Path at which to create cleanfreak directory (Default is Home).")
}
