package cmd

import (
	"fmt"
	"os"
	"path"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Runs cleanfreak on the downloads folder.",
	Long: `This command will clean the user's downloads folder by prompting the user
	to either transfer those files to the appropriate location in the cleanfreak project directory,
	or remove them.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Always relative to home!
		homedir, err := os.UserHomeDir() // Add error handling.
		cf_root := viper.Get("directory")

		if err != nil {
			panic(err)
		}

		str, ok := cf_root.(string)
		if !ok {
			fmt.Println("\n🚨 Error - could not find cf_root in config.")
			return
		}

		// Check whether diagnose passes.
		if core.DirEmpty("Downloads") {
			cmdutil.PrintDiagnoseSuccess("No files in the Downloads folder.")
			fmt.Println("\nEverything is in order! 🎉")
			return
		}

		// If not, initialise cleanfreak process.
		cmdutil.PrintDiagnoseFail("No files in the Downloads folder.")

		// Does project directory exist?
		if !core.DirExists(path.Join(homedir, str)) {
			fmt.Println("\n🚨 Error - could not find cleanfreak project directory. Kindly execute cf init before running cf clean.")
			return
		}

	},
}

func init() {
	cleanCmd.AddCommand(downloadsCmd)
}
