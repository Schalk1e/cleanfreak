package cmd

import (
	"log"
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
	var wg sync.WaitGroup

	wg.Add(len(cache_dirs))

	for _, dir := range cache_dirs {
		go func(dir string) {
			defer wg.Done()
			size, _ := core.DirSize(dir)
			data := []string{
				dir,
				cmdutil.ByteStringParse(strconv.FormatInt(size, 10)),
			}
			cache_data = append(cache_data, data)
		}(dir)
	}

	wg.Wait()

	data := cmdutil.FilterSlice(cache_data, threshold)

	if len(data) > 0 {
		cmdutil.PrintDiagnoseFail(diagnose_text)
		for _, d := range data {
			files := core.List(d[0], false)
			action := cmdutil.CacheDeleteSurvey(d[0], d[1])

			c := core.Clean{}
			if action == "Y" {
				for _, file := range files[1:] {
					c.SourceFile = file
					if err := c.FileDelete(); err != nil {
						log.Printf("Failed to delete file: %v", err)
					}
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
