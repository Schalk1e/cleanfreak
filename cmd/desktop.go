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

var desktopCmd = &cobra.Command{
	Use:   "desktop",
	Short: "Runs cleanfreak on the desktop folder.",
	Long: `This command will clean the user's desktop folder by prompting the user
	to either transfer those files to the appropriate location in the cleanfreak project directory,
	or remove them.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Always relative to home! Think abt this a little more carefully.
		homedir, err := os.UserHomeDir() // Add error handling.
		cf_root := viper.Get("directory")
		diagnose_text := "No files in the Desktop folder."

		if err != nil {
			panic(err)
		}

		str, ok := cf_root.(string)
		if !ok {
			fmt.Println("Could not find cf_root in config.")
			return
		}

		if !core.DirExists(path.Join(homedir, str)) {
			fmt.Println("\n Could not find cleanfreak project directory. Kindly execute cf init before running cf clean.")
			return
		}

		if core.DirEmpty("Desktop") {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			fmt.Println("\nEverything is in order! 🎉")
			return
		}

		cmdutil.PrintDiagnoseFail(diagnose_text)

		CleanDesktop(str)

	},
}

func CleanDesktop(target string) {

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var files []string
	desktop := path.Join(homedir, "Desktop")

	files = core.List(desktop, false)

	for _, file := range files[1:] {

		filename := path.Base(file)
		c := core.Clean{}

		action := cmdutil.FileSurvey(filename)

		if action == "Move" {
			c.SourceFile = file
			c.TargetFile = core.DirSurvey(path.Join(homedir, target))
			c.FileTransfer()
		} else if action == "Delete" {
			c.SourceFile = file
			c.FileDelete()
		} else {
			break
		}
	}
}

func init() {
	cleanCmd.AddCommand(desktopCmd)
}
