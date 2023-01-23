package cmdutil

import (
	"strconv"
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
