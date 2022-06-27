//go:build !windows

package cmdutil

import (
	"fmt"
)

func PrintDiagnoseSuccess(msg string) {
	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰ " + green + tick + " Success" + end)
}

func PrintDiagnoseFail(msg string) {
	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰ " + red + ballot + " Failed" + end)
}

func PrintArrows(msgs []string) {
	for i, msg := range msgs {
		if i == 0 {
			fmt.Println("\n" + blue + arrow + " " + end + msg)
		} else {
			fmt.Println(blue + arrow + " " + end + msg)
		}
	}
}

func PrintMoved() {
	fmt.Println("╰ " + green + tick + " Moved" + end)
}

func PrintDeleted() {
	fmt.Println("╰ " + green + tick + " Removed" + end)
}

func PrintWarning(msg string) {
	fmt.Println("\n" + bold + red + "‼" + end + " " + msg)
}

func PrintOrder() {
	fmt.Println("\nEverything is in order! " + celebrate)
}
