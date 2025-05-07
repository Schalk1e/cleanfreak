//go:build !windows

package cmd

import (
	_ "embed"
	"fmt"
	"os"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed embed/banner.txt
var banner string

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "Welcome to Cleanfreak! An opinionated workspace organisation and cleaning utility.",
	Long:  banner,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cleanfreak.yaml)")
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
