package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Runs cleanfreak on the downloads folder.",
	Long: `
This command will clean the user's downloads folder by prompting the user
to either transfer those files to the appropriate location in the cleanfreak
project directory, or remove them.`,

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
			fmt.Println("\nCould not find cleanfreak project directory. Kindly execute cf init before running cf clean.")
			return
		}

		d := core.Dir{}
		downloadsdir := d.GetDownloads()
		if !core.DirExists(downloadsdir) {
			cmdutil.PrintDirectoryNotFound(downloadsdir)
		}
		if core.DirEmpty(d.GetDownloads()) {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			cmdutil.PrintOrder()
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

	downloads := path.Join(homedir, "Downloads")
	files := core.List(downloads, false)

	for i := 1; i < len(files); i++ {

		filename := path.Base(files[i])
		c := core.Clean{}
		action := cmdutil.FileSurvey(filename)

		if action == "Move" {

			opts := core.List(path.Join(homedir, target), true)
			folder := cmdutil.DirSurvey(opts)
			name := cmdutil.FileNameSurvey()

			c.SourceFile = files[i]
			c.TargetFile = path.Join(folder, name)
			if err := c.FileTransfer(); err != nil {
				log.Printf("File not successfully transfered: %v", err)
			}
			cmdutil.PrintMoved()

		} else if action == "Delete" {

			c.SourceFile = files[i]

			if err := c.FileDelete(); err != nil {
				log.Printf("File not successfully removed: %v", err)
			}

			cmdutil.PrintDeleted()

		} else if action == "View" {

			open_err := open.Run(files[i])
			if open_err != nil {
				fmt.Println(open_err.Error())
			}
			i = i - 1

		} else {
			break
		}
	}
}

func init() {
	cleanCmd.AddCommand(downloadsCmd)
}
