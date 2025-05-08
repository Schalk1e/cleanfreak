package clean

import (
	"github.com/spf13/cobra"
)

var AllCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs cleanfreak across all items.",
	Long: `This command will execute the cleanfreak process over all items
	listed by cleanfreak list items.`,
	Run: func(cmd *cobra.Command, args []string) {
		DesktopCmd.Run(cmd, args)
		DownloadsCmd.Run(cmd, args)
		TrashCmd.Run(cmd, args)
		CacheCmd.Run(cmd, args)
	},
}
