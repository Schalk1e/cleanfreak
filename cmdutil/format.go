package cmdutil

import (
	"fmt"
)

func PrintDiagnoseSuccess(msg string) {
	var bold = Bold()
	var green = Green()
	var cyan = Cyan()
	var end = End()

	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰-" + green + "✔ Success" + end)
}

func PrintDiagnoseFail(msg string) {
	var bold = Bold()
	var cyan = Cyan()
	var red = Red()
	var end = End()

	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰-" + red + "✘ Failed" + end)
}

func PrintArrows(msgs []string) {
	var blue = Cyan()
	var end = End()

	for i, msg := range msgs {
		if i == 0 {
			fmt.Println("\n" + blue + "⮕ " + end + msg)
		} else {
			fmt.Println(blue + "⮕ " + end + msg)
		}
	}
}
