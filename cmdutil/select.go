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
