//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "Welcome to Cleanfreak! An opinionated workspace organisation and cleaning utility.",
	Long: `

   ______________                  ________                  ______  
   __  ____/__  /__________ __________  __/_________________ ___  /__
   _  /    __  /_  _ \  __ '/_  __ \_  /_ __  ___/  _ \  __ '/_  //_/
   / /___  _  / /  __/ /_/ /_  / / /  __/ _  /   /  __/ /_/ /_  ,<   
   \____/  /_/  \___/\__,_/ /_/ /_//_/    /_/    \___/\__,_/ /_/|_|  
																	 

Welcome to Cleanfreak! An opinionated workspace organisation and cleaning 
utility. 

$ cf list items`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(cmd.initConfig)
	core.EnableVT()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cleanfreak.yaml)")
}
