package cmdutil

import (
	"fmt"

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

func DeleteSurvey() string {

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
