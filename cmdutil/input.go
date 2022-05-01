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
					"Stop",
				},
			},
		},
	}

	action := struct {
		Action string
	}{}

	fmt.Println("")
	survey.Ask(actionSurvey, &action)

	return action.Action
}

func DirSurvey(opts []string) string {

	dirSurvey := []*survey.Question{
		{
			Name: "dir",
			Prompt: &survey.Select{
				Message: fmt.Sprintf("Select target:"),
				Options: opts,
			},
		},
	}

	target := struct {
		Dir string
	}{}

	survey.Ask(dirSurvey, &target)

	return target.Dir
}

func FileNameSurvey() string {

	name := ""
	prompt := &survey.Input{
		Message: "Please enter filename: ",
	}

	survey.AskOne(prompt, &name)

	return name
}
