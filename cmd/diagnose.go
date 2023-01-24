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
		threshold := viper.GetFloat64("threshold")
		cacheReport(cache_dirs, threshold)
	},
}

func cacheReport(cache_dirs []string, threshold float64) {
	var cache_data [][]string
	num := len(cache_dirs)
	sum := 0.0

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

	data := cmdutil.FilterSlice(cache_data, threshold)

	if len(data) > 0 {
		cmdutil.PrintDiagnoseFail("No caches over the configured threshold.")
		fmt.Println("")
		fmt.Println(
			cmdutil.TableFromSlices(
				cmdutil.OrderSliceByFloat(
					data,
				),
			),
		)
		for i := 0; i < len(cache_data); i++ {
			sum += cmdutil.FloatFromGBString(
				cache_data[i][len(cache_data[i])-1],
			)
		}
		cmdutil.PrintCacheTotal((strconv.FormatFloat(sum, 'f', 2, 64) + "GB"))
	} else {
		cmdutil.PrintDiagnoseSuccess("No caches over the configured threshold.")
	}
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}
