package clean

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

var DesktopCmd = &cobra.Command{
	Use:   "desktop",
	Short: "Runs cleanfreak on the desktop folder.",
	Long: `
This command will clean the user's desktop folder by prompting the user
to either transfer those files to the appropriate location in the cleanfreak
project directory, or remove them.`,

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
			fmt.Println("\nCould not find cleanfreak project directory. Kindly execute cf init before running cf clean.")
			return
		}

		d := core.Dir{}
		desktopdir := d.GetDesktop()
		if !core.DirExists(desktopdir) {
			cmdutil.PrintDirectoryNotFound(desktopdir)
			return
		}
		if core.DirEmpty(desktopdir) {
			cmdutil.PrintDiagnoseSuccess(diagnose_text)
			cmdutil.PrintOrder()
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

	desktop := path.Join(homedir, "Desktop")
	files := core.List(desktop, false)

	for i := 1; i < len(files); i++ {

		filename := path.Base(files[i])
		c := core.Clean{}
		action := cmdutil.FileSurvey(filename)

		switch action {
		case "Move":
			opts := core.List(path.Join(homedir, target), true)
			folder := cmdutil.DirSurvey(opts)
			name := cmdutil.FileNameSurvey(filename)

			c.SourceFile = files[i]
			c.TargetFile = path.Join(folder, name)
			if err := c.FileTransfer(); err != nil {
				log.Printf("File not successfully transfered: %v", err)
			}

			cmdutil.PrintMoved()
		case "Delete":
			c.SourceFile = files[i]
			if err := c.FileDelete(); err != nil {
				log.Printf("File not successfully removed: %v", err)
			}

			cmdutil.PrintDeleted()
		case "View":
			open_err := open.Run(files[i])
			if open_err != nil {
				fmt.Println(open_err.Error())
			}
			i = i - 1
		default:
			break
		}
	}
}
