package cmd

import (
	"fmt"
	"os"
	"path"

	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
)

var rerun bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a cleanfreak folder.",
	Long: `This command will initialise a default cleanfreak folder in the specified
	location that contains a number of subfolders that are intended to provide appropriate homes
	for most filetypes.`,
	Run: func(cmd *cobra.Command, args []string) {
		cf_root := "cleanfreak" // Get this from config...
		base_dir, _ := cmd.Flags().GetString("path")
		rerun, _ := cmd.Flags().GetBool("rerun")

		if core.DirExists(cf_root) && !rerun {
			response := fmt.Sprintf("Already contains a directory named %s!", cf_root)
			fmt.Println(response)
			os.Exit(1)
		}

		subdirs := []string{"personal", "work"} // Get this from config...
		if base_dir == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			} else {
				core.DirsAdd(path.Join(homedir, cf_root), subdirs)
			}
		} else {
			core.DirsAdd(path.Join(base_dir, cf_root), subdirs)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("path", "", "Path at which to create cleanfreak directory (Default is Home).")
	initCmd.Flags().BoolVar(&rerun, "rerun", false, "Re-init cleanfreak directory from config - adds new directories and keeps existing.")
}
