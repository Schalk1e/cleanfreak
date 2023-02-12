//go:build windows

package cmdutil

import (
	"fmt"
	"strings"
)

func PrintDiagnoseSuccess(msg string) {
	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰ " + green + "Success" + end)
}

func PrintDiagnoseFail(msg string) {
	fmt.Println("\n" + bold + cyan + "Done: " + end + msg)
	fmt.Println("╰ " + red + "Failed" + end)
}

func PrintCacheTotal(total string) {
	fmt.Println(bold + cyan + "Total Storage: " + end + " " + total)
}

func PrintArrows(msgs []string) {
	for i, msg := range msgs {
		if i == 0 {
			fmt.Println("\n" + blue + "- " + end + msg)
		} else {
			fmt.Println(blue + "- " + end + msg)
		}
	}
}

func PrintMoved() {
	fmt.Println("╰ " + green + "Moved" + end)
}

func PrintDeleted() {
	fmt.Println("╰ " + green + "Removed" + end)
}

func PrintWarning(msg string) {
	fmt.Println("\n" + bold + red + "‼" + end + " " + msg)
}

func PrintOrder() {
	fmt.Println("\nEverything is in order!")
}
