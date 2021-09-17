package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check compliance of workspace.",
	Long: `
The diagnose command will check whether the host workspace complies with each of the 10 guiding
principles of a clean workspace without taking any explicit action. The results will be reported
back to the user.`,
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diagnose called")
	},
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
