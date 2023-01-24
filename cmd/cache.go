package cmd

import (
	"strconv"
	"sync"

	cmdutil "github.com/Schalk1e/cleanfreak/cmdutil"
	"github.com/Schalk1e/cleanfreak/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Runs cleanfreak on the user's cache folder.",
	Long:  `This command will clean the user's cache folder.`,

	Run: func(cmd *cobra.Command, args []string) {
		diagnose_text := "No caches over the configured threshold."
		cache_dirs := viper.GetStringSlice("cachedirs")
		threshold := viper.GetFloat64("threshold")

		CleanCache(
			cache_dirs,
			threshold,
			diagnose_text,
		)
	},
}

func CleanCache(cache_dirs []string, threshold float64, diagnose_text string) {
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

	data := cmdutil.FilterSlice(cache_data, threshold)

	if len(data) > 0 {
		cmdutil.PrintDiagnoseFail(diagnose_text)
		for i := 0; i < len(data); i++ {
			files := core.List(data[i][0], false)
			action := cmdutil.CacheDeleteSurvey(
				data[i][0],
				data[i][1],
			)
			c := core.Clean{}
			if action == "Y" {
				for i := 1; i < len(files); i++ {
					c.SourceFile = files[i]
					c.FileDelete()
				}
				cmdutil.PrintCleaned()
			}
		}
	} else {
		cmdutil.PrintDiagnoseSuccess(diagnose_text)
		cmdutil.PrintOrder()
	}
}

func init() {
	cleanCmd.AddCommand(cacheCmd)
}
