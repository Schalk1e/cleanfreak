package core

import (
	"reflect"
	"testing"
)

// ShowTestResult compares the expected and actual results, printing an error message if they differ.
func ShowTestResult(got, want any, t *testing.T) {
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

// Use cmp.Equal instead?
func ShowTestResultDeepEqual(got, want any, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}
