package cmdutil

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func FileSurvey(filename string) string {

	actionSurvey := []*survey.Question{
		{
			Name: "action",
			Prompt: &survey.Select{
				Message: fmt.Sprintf("Select action to take on file %s:", filename),
				Options: []string{
					"Move",
					"Delete",
					"View",
					"Stop",
				},
			},
		},
	}

	action := struct {
		Action string
	}{}

	fmt.Println("")
	survey_err := survey.Ask(actionSurvey, &action)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return action.Action
}

func DirSurvey(opts []string) string {

	dirSurvey := []*survey.Question{
		{
			Name: "dir",
			Prompt: &survey.Select{
				Message: fmt.Sprintln("Select target:"),
				Options: opts,
			},
		},
	}

	target := struct {
		Dir string
	}{}

	survey_err := survey.Ask(dirSurvey, &target)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return target.Dir
}

func FileNameSurvey() string {

	name := ""
	prompt := &survey.Input{
		Message: "Please enter filename: ",
	}

	survey_err := survey.AskOne(prompt, &name)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return name
}

func TrashDeleteSurvey() string {

	action := ""
	prompt := &survey.Input{
		Message: "Would you like to clear out the trash? (Y/N): ",
	}

	fmt.Println("")
	survey_err := survey.AskOne(prompt, &action)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return action
}

func CacheDeleteSurvey(path string, size string) string {
	action := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("Would you like to clean cache directory %s of size %s? (Y/N):",
			path,
			size,
		),
	}

	fmt.Println("")
	survey_err := survey.AskOne(prompt, &action)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return action
}

func ConfigRoot() string {
	name := ""
	prompt := &survey.Input{
		Message: "Please enter the path where the cleanfreak directory should be created: ",
	}

	survey_err := survey.AskOne(prompt, &name)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return name

}

func ConfigInitPaths() []string {
	input := ""

	prompt := &survey.Input{
		Message: "Please enter the folder names to initialise the cleanfreak directory separated with comma's (example: one,two): ",
	}

	survey_err := survey.AskOne(prompt, &input)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}
	init_paths := strings.Split(input, ",")
	for i := range init_paths {
		init_paths[i] = strings.TrimSpace(init_paths[i])
	}

	return init_paths
}

func CachePaths() []string {
	input := ""

	prompt := &survey.Input{
		Message: "Please enter the full cache paths you'd like cleafreak to watch separated with comma's (example: /path/one,/path/two): ",
	}

	survey_err := survey.AskOne(prompt, &input)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}
	cache_paths := strings.Split(input, ",")
	for i := range cache_paths {
		cache_paths[i] = strings.TrimSpace(cache_paths[i])
	}

	return cache_paths
}

func CacheSize() float64 {
	var size float64
	prompt := &survey.Input{
		Message: "Please enter the allowed threshold for cache size in GB's: ",
	}

	survey_err := survey.AskOne(prompt, &size)
	if survey_err != nil {
		fmt.Println(survey_err.Error())
	}

	return size
}
