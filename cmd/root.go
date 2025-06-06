//go:build !windows

package cmd

import (
	"fmt"
	"os"

	"github.com/Schalk1e/cleanfreak/cmd/clean"
	"github.com/Schalk1e/cleanfreak/cmd/plan"
	"github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var banner string

var cfgFile string
var RootCmd = &cobra.Command{
	Use:   "cf",
	Short: "A basic workspace organisation utility.",
	Long:  banner,
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cleanfreak.yaml)")
	RootCmd.AddCommand(diagnoseCmd)
	RootCmd.AddCommand(listCmd)
	listCmd.AddCommand(itemsCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(clean.CleanCmd)
	clean.CleanCmd.AddCommand(clean.AllCmd)
	clean.CleanCmd.AddCommand(clean.DownloadsCmd)
	clean.CleanCmd.AddCommand(clean.DesktopCmd)
	clean.CleanCmd.AddCommand(clean.CacheCmd)
	clean.CleanCmd.AddCommand(clean.TrashCmd)
	RootCmd.AddCommand(plan.PlanCmd)
	plan.PlanCmd.AddCommand(plan.DownloadsCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cleanfreak")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		fmt.Println("Creating default config file in home directory.")

		cmdutil.BuildConfig(home)
		viper_err := viper.ReadInConfig()
		if viper_err != nil {
			fmt.Println(viper_err.Error())
		}
	}
}
