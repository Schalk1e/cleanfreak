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
