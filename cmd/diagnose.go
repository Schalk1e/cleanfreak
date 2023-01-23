package cmd

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	core "github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check compliance of workspace.",
	Long: `
The diagnose command will check whether the host workspace complies with each of the items
shown by cf list items without taking any explicit action. The results will be reported
back to the user.`,

	Run: func(cmd *cobra.Command, args []string) {
		d := core.Dir{}
		d.OS = runtime.GOOS

		dirs := []string{d.GetDownloads(), d.GetDesktop(), d.GetTrash()}
		dirnames := []string{"Downloads", "Desktop", "Trash"}

		trash_supported := cmdutil.IsIn(d.OS, cmdutil.TrashSupported[:])

		if !trash_supported {
			dirs = dirs[:2]
			dirnames = dirnames[:2]
		}
		for indx, dir := range dirs {
			if core.DirEmpty((dir)) {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseSuccess(text)
			} else {
				text := fmt.Sprintf("No files in the %s folder.", dirnames[indx])
				cmdutil.PrintDiagnoseFail(text)
			}
		}
		if !trash_supported {
			cmdutil.PrintWarning("Cleaning Trash not yet supported on this OS! Please contribute!")
		}
		cache_dirs := viper.GetStringSlice("cachedirs")
		cacheReport(cache_dirs)
	},
}

func cacheReport(cache_dirs []string) {
	var cache_data [][]string
	num := len(cache_dirs)

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			size, _ := core.DirSize(cache_dirs[i])
			data := []string{
				cache_dirs[i],
				cmdutil.ByteStringParse(strconv.FormatInt(size, 10)),
			}
			cache_data = append(cache_data, data)
		}(i)
	}

	wg.Wait()

	fmt.Println("")
	fmt.Println(
		cmdutil.PrintTableFromSlices(
			cmdutil.OrderSliceByFloat(cache_data),
		),
	)
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
