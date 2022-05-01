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
		// Always relative to home! Think abt this a little more carefully.
		homedir, err := os.UserHomeDir() // Add error handling.
		cf_root := viper.Get("directory")
		diagnose_text := "No files in the Downloads folder."

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

		if core.DirEmpty("Downloads") {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			fmt.Println("\nEverything is in order! 🎉")
			return
		}

		cmdutil.PrintDiagnoseFail(diagnose_text)

		CleanDownloads(str)

	},
}

func CleanDownloads(target string) {

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var files []string
	downloads := path.Join(homedir, "Downloads")

	files = core.List(downloads, false)

	for _, file := range files[1:] {

		filename := path.Base(file)
		c := core.Clean{}

		action := cmdutil.FileSurvey(filename)

		if action == "Move" {

			opts := core.List(path.Join(homedir, target), true)

			folder := cmdutil.DirSurvey(opts)
			name := cmdutil.FileNameSurvey()

			c.SourceFile = file
			c.TargetFile = path.Join(folder, name)
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
	cleanCmd.AddCommand(downloadsCmd)
}
