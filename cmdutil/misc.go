package cmdutil

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// IsIn checks whether a given string exists in a slice of strings.
// Consider using `slices.Contains(list, a)` from `golang.org/x/exp/slices`.
func IsIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringDifference finds string elements that are present in slice a
// but not present in slice b. It does so by pre-allocating a map with
// keys set to the strings in b. This creates a lookup to check the
// elements of a against.
func StringDifference(a []string, b []string) []string {
	InB := make(map[string]bool, len(b))
	for _, s := range b {
		InB[s] = true
	}

	difference := []string{}
	for _, s := range a {
		if !InB[s] {
			difference = append(difference, s)
		}
	}

	return difference
}

// ByteStringParse converts a numeric string representing bytes into a human-readable gigabyte (GB) string.
// Example: "1100000000" -> "1.10GB"
func ByteStringParse(decimal_byte_str string) string {
	int_value, err := strconv.Atoi(decimal_byte_str)
	if err != nil {
		return "Invalid Input"
	}

	decimal_value := float64(int_value) / 1_000_000_000
	return fmt.Sprintf("%.2fGB", decimal_value)
}

// FloatFromGBString extracts and converts the numeric portion of a string formatted in "GB" to a float.
// Example: "10.5GB" -> 10.5
func FloatFromGBString(gb_string string) float64 {
	num := strings.TrimSuffix(gb_string, "GB")
	num_float, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0 // Returning 0 as default for invalid input
	}
	return num_float
}

// OrderSliceByFloat sorts a 2D string slice in descending order based on the float value in the last column.
func OrderSliceByFloat(input [][]string) [][]string {
	floatValues := make([]float64, len(input))
	for i, row := range input {
		floatValues[i] = FloatFromGBString(row[len(row)-1])
	}

	sort.SliceStable(input, func(i, j int) bool {
		return floatValues[i] > floatValues[j]
	})

	return input
}

// FilterSlice filters out rows from a 2D slice where the last column's float value is below a threshold.
// Example: If threshold = 6.0, only rows where the last column > 6.0 are retained.
func FilterSlice(input [][]string, threshold float64) [][]string {
	filtered_slice := [][]string{}

	for _, item := range input {
		if FloatFromGBString(item[len(item)-1]) > threshold {
			filtered_slice = append(filtered_slice, item)
		}
	}

	return filtered_slice
}

// TableFromSlices generates a formatted table string representation of a 2D slice of strings.
// Example:
// +----+----+
// | A  | B  |
// | C  | D  |
// +----+----+
func TableFromSlices(input [][]string) string {
	if len(input) == 0 {
		return ""
	}

	// Determine max column widths
	lengths := make([]int, len(input[0]))
	for _, row := range input {
		for i, cell := range row {
			if len(cell) > lengths[i] {
				lengths[i] = len(cell)
			}
		}
	}

	// Construct separator line
	sep_line := "+"
	for _, col_len := range lengths {
		sep_line += strings.Repeat("-", col_len+2) + "+"
	}

	var builder strings.Builder
	builder.WriteString(sep_line + "\n")
	for _, row := range input {
		builder.WriteString("|")
		for i, cell := range row {
			builder.WriteString(" " + cell + strings.Repeat(" ", lengths[i]-len(cell)) + " |")
		}
		builder.WriteString("\n" + sep_line + "\n")
	}

	return builder.String()
}
