package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cleanfreak",
	Short: "Welcome to Cleanfreak! An opinionated workspace organisation and cleaning utility.",
	Long: `

   ______________                  ________                  ______  
   __  ____/__  /__________ __________  __/_________________ ___  /__
   _  /    __  /_  _ \  __ '/_  __ \_  /_ __  ___/  _ \  __ '/_  //_/
   / /___  _  / /  __/ /_/ /_  / / /  __/ _  /   /  __/ /_/ /_  ,<   
   \____/  /_/  \___/\__,_/ /_/ /_//_/    /_/    \___/\__,_/ /_/|_|  
																	 

Welcome to Cleanfreak! An opinionated workspace organisation and cleaning 
utility. We ascribe to 10 guiding principles of a clean workspace. To see 
them run:

$ cleanfreak list principles 

We can't get the trash out of the ocean - let's get it off our PC's! ✨🖥️✨`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// global flags & config.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cleanfreak.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cleanfreak" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cleanfreak")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}