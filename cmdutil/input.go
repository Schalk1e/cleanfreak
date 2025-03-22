// Need to migrate this dependency to https://github.com/charmbracelet/bubbletea

package cmdutil

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// FileSurvey prompts the user to select an action to take on the specified file.
//
// This function presents a list of actions (Move, Delete, View, Stop) and returns the user's choice.
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
	if surveyErr := survey.Ask(actionSurvey, &action); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return action.Action
}

// DirSurvey prompts the user to select a target directory from the provided options.
//
// This function displays a list of directories and returns the user's choice.
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

	if surveyErr := survey.Ask(dirSurvey, &target); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return target.Dir
}

// FileNameSurvey prompts the user to input a filename.
//
// This function simply asks the user to enter a filename and returns the input.
func FileNameSurvey() string {
	name := ""
	prompt := &survey.Input{
		Message: "Please enter filename: ",
	}

	if surveyErr := survey.AskOne(prompt, &name); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return name
}

// TrashDeleteSurvey prompts the user if they would like to clear out the trash (Y/N).
//
// The function expects a yes/no answer and returns the response.
func TrashDeleteSurvey() string {
	action := ""
	prompt := &survey.Input{
		Message: "Would you like to clear out the trash? (Y/N): ",
	}

	fmt.Println("")
	if surveyErr := survey.AskOne(prompt, &action); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return action
}

// CacheDeleteSurvey prompts the user if they want to clean a specific cache directory.
//
// This function accepts a path and size for the cache directory, then asks the user for confirmation (Y/N).
func CacheDeleteSurvey(path string, size string) string {
	action := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("Would you like to clean cache directory %s of size %s? (Y/N):", path, size),
	}

	fmt.Println("")
	if surveyErr := survey.AskOne(prompt, &action); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return action
}

// ConfigRoot prompts the user to enter the path where the cleanfreak directory should be created.
//
// This function returns the user's input for the directory path.
func ConfigRoot() string {
	name := ""
	prompt := &survey.Input{
		Message: "Please enter the path where the cleanfreak directory should be created: ",
	}

	if surveyErr := survey.AskOne(prompt, &name); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return name
}

// ConfigInitPaths prompts the user to input a comma-separated list of folder names to initialize.
//
// The function returns the list of folder names as a slice of strings.
func ConfigInitPaths() []string {
	input := ""

	prompt := &survey.Input{
		Message: "Please enter the folder names to initialise the cleanfreak directory separated with comma's (example: one,two): ",
	}

	if surveyErr := survey.AskOne(prompt, &input); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	initPaths := strings.Split(input, ",")
	for i := range initPaths {
		initPaths[i] = strings.TrimSpace(initPaths[i])
	}

	return initPaths
}

// CachePaths prompts the user to input a comma-separated list of cache paths to watch.
//
// The function returns a slice of cache paths entered by the user.
func CachePaths() []string {
	input := ""

	prompt := &survey.Input{
		Message: "Please enter the full cache paths you'd like cleanfreak to watch separated with comma's (example: /path/one,/path/two): ",
	}

	if surveyErr := survey.AskOne(prompt, &input); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	cachePaths := strings.Split(input, ",")
	for i := range cachePaths {
		cachePaths[i] = strings.TrimSpace(cachePaths[i])
	}

	return cachePaths
}

// CacheSize prompts the user to enter the allowed threshold for cache size in GB.
//
// This function returns the entered cache size as a float64 value.
func CacheSize() float64 {
	var size float64
	prompt := &survey.Input{
		Message: "Please enter the allowed threshold for cache size in GB's: ",
	}

	if surveyErr := survey.AskOne(prompt, &size); surveyErr != nil {
		fmt.Println(surveyErr.Error())
	}

	return size
}
