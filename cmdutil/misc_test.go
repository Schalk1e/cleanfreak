package cmdutil

import (
	"testing"

	core "github.com/Schalk1e/cleanfreak/core"
)

// TestIsIn verifies whether the IsIn function correctly determines
// if a given string is present in a slice of strings.
func TestIsIn(t *testing.T) {
	cases := []struct {
		name        string
		search_str  string
		search_list []string
		result      bool
	}{
		{"string_is_in", "a", []string{"a", "b", "c"}, true},
		{"string_not_in", "a", []string{"b", "c", "d"}, false},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := IsIn(testcase.search_str, testcase.search_list)
			want := testcase.result

			core.ShowTestResult(got, want, t)
		})
	}
}

// TestByteStringParse ensures that ByteStringParse correctly formats
// byte values as human-readable GB strings.
func TestByteStringParse(t *testing.T) {
	cases := []struct {
		name                string
		decimal_byte_string string
		result              string
	}{
		{"more_than_1gb", "1100000000", "1.10GB"},
		{"less_than_1gb", "900000000", "0.90GB"},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := ByteStringParse(testcase.decimal_byte_string)
			want := testcase.result

			core.ShowTestResult(got, want, t)
		})
	}
}

// TestFloatFromGBString verifies that FloatFromGBString correctly
// extracts the float value from a formatted GB string.
func TestFloatFromGBString(t *testing.T) {
	cases := []struct {
		name      string
		gb_string string
		result    float64
	}{
		{"single_digit", "1GB", 1},
		{"double_digit", "0.90GB", 0.9},
		{"large_value", "10.00GB", 10},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := FloatFromGBString(testcase.gb_string)
			want := testcase.result

			core.ShowTestResult(got, want, t)
		})
	}
}

// TestOrderSliceByFloat checks that OrderSliceByFloat correctly sorts
// a 2D string slice based on the float values in the last column.
func TestOrderSliceByFloat(t *testing.T) {
	tests := []struct {
		input  [][]string
		output [][]string
	}{
		{
			input: [][]string{
				{"A", "B", "10.5"},
				{"C", "D", "5.3"},
				{"E", "F", "7.2"},
			},
			output: [][]string{
				{"A", "B", "10.5"},
				{"E", "F", "7.2"},
				{"C", "D", "5.3"},
			},
		},
		{
			input: [][]string{
				{"A", "B", "1.0"},
				{"C", "D", "2.0"},
				{"E", "F", "3.0"},
			},
			output: [][]string{
				{"E", "F", "3.0"},
				{"C", "D", "2.0"},
				{"A", "B", "1.0"},
			},
		},
	}

	for _, testcase := range tests {
		t.Run("OrderSliceByFloat", func(t *testing.T) {
			got := OrderSliceByFloat(testcase.input)
			want := testcase.output

			core.ShowTestResultDeepEqual(got, want, t)
		})
	}
}

// TestFilterSlice ensures that FilterSlice correctly filters out
// rows where the last column’s float value is below a given threshold.
func TestFilterSlice(t *testing.T) {
	tests := []struct {
		input     [][]string
		threshold float64
		output    [][]string
	}{
		{
			input: [][]string{
				{"A", "B", "10.5"},
				{"C", "D", "5.3"},
				{"E", "F", "7.2"},
			},
			threshold: 6.0,
			output: [][]string{
				{"A", "B", "10.5"},
				{"E", "F", "7.2"},
			},
		},
		{
			input: [][]string{
				{"A", "B", "1.0"},
				{"C", "D", "2.0"},
				{"E", "F", "3.0"},
			},
			threshold: 2.5,
			output: [][]string{
				{"E", "F", "3.0"},
			},
		},
	}

	for _, testcase := range tests {
		t.Run("FilterSlice", func(t *testing.T) {
			got := FilterSlice(testcase.input, testcase.threshold)
			want := testcase.output

			core.ShowTestResultDeepEqual(got, want, t)
		})
	}
}
