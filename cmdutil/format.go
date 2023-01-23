//go:build !windows

package cmdutil

import (
	"fmt"
	"strings"
)

func PrintTableFromSlices(input [][]string) string {
	lengths := make([]int, len(input[0]))
	for _, row := range input {
		for i, cell := range row {
			if len(cell) > lengths[i] {
				lengths[i] = len(cell)
			}
		}
	}
	sep_line := "+"
	for _, col_len := range lengths {
		sep_line += strings.Repeat("-", col_len+2) + "+"
	}
	table := sep_line + "\n"
	for _, row := range input {
		table += "|"
		for i, cell := range row {
			table += " " + cell + strings.Repeat(" ", lengths[i]-len(cell)) + " |"
		}
		table += "\n" + sep_line + "\n"
	}

	return table
}

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
	fmt.Println("\n" + green + tick + " Everything is in order! No action to take." + end)
}
