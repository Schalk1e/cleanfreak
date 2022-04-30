package cmd

import (
	"fmt"
	"os"
	"path"

	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rerun bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a cleanfreak project folder.",
	Long: `This command will initialise a default cleanfreak folder in the specified
	location that contains a number of subfolders that are intended to provide appropriate homes
	for most filetypes.`,
	Run: func(cmd *cobra.Command, args []string) {
		cf_root := viper.Get("directory")
		base_dir, _ := cmd.Flags().GetString("path")
		rerun, _ := cmd.Flags().GetBool("rerun")
		str, ok := cf_root.(string)
		if !ok {
			panic("Could not find cf_root in config.")
		}

		if core.DirExists(str) && !rerun {
			response := fmt.Sprintf("Already contains a directory named %s!", cf_root)
			fmt.Println(response)
			os.Exit(1)
		}

		subdirs := viper.GetStringSlice("subdirs")
		if base_dir == "" {
			homedir, err := os.UserHomeDir()

			if err != nil {
				panic(err)
			} else {
				core.DirsAdd(path.Join(homedir, str), subdirs)
			}
		} else {
			core.DirsAdd(path.Join(base_dir, str), subdirs)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("path", "", "Path at which to create cleanfreak directory (Default is Home).")
	initCmd.Flags().BoolVar(&rerun, "rerun", false, "Re-init cleanfreak directory from config - adds new directories and keeps existing.")
}
