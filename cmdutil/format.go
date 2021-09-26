package cmdutil

import (
	"fmt"
)

func PrintDiagnoseSuccess() {
	var bold = Bold()
	var green = Green()
	var cyan = Cyan()
	var end = End()

	fmt.Println("\n" + bold + cyan + "Done: " + end + "No files in the Downloads folder.")
	fmt.Println("╰-" + green + "✔ Success" + end)
}

func PrintDiagnoseFail() {
	var bold = Bold()
	var cyan = Cyan()
	var red = Red()
	var end = End()

	fmt.Println("\n" + bold + cyan + "Done: " + end + "No files in the Downloads folder.")
	fmt.Println("╰-" + red + "✘ Failed" + end)
}
