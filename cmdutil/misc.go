package cmdutil

import (
	"sort"
	"strconv"
	"strings"
)

func IsIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ByteStringParse(decimal_byte_str string) string {
	int_value, _ := strconv.Atoi(decimal_byte_str)

	decimal_value := float64(int_value) / float64(1000000000)

	return strconv.FormatFloat(decimal_value, 'f', 2, 64) + "GB"
}

func FloatFromGBString(gb_string string) float64 {
	num := strings.Replace(gb_string, "GB", "", 1)

	num_float, _ := strconv.ParseFloat(num, 64)

	return num_float
}

func OrderSliceByFloat(input [][]string) [][]string {
	sort.Slice(input, func(i, j int) bool {
		return FloatFromGBString(
			input[i][len(input[i])-1],
		) >
			FloatFromGBString(
				input[j][len(input[j])-1],
			)
	})
	return input
}

func FilterSlice(input [][]string, threshold float64) [][]string {
	filtered_slice := [][]string{}

	for i := 0; i < len(input); i++ {
		if FloatFromGBString(input[i][len(input[i])-1]) > threshold {
			filtered_slice = append(filtered_slice, input[i])
		}
	}

	return filtered_slice
}

func TableFromSlices(input [][]string) string {
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
