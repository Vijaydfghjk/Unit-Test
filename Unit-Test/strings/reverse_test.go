package strings

import (
	"testing"
)

func TestReverse(t *testing.T) {

	tests := []struct {
		val string

		expected string
	}{

		{"Vijay", "yajiV"},
		{"Ram", "maR"},
		{"Vi", "iV"},
		{"", ""},
	}

	for _, tt := range tests {

		result := Reverse(tt.val)

		if result != tt.expected {

			t.Errorf("Reverse(%s)=%s; want%s", tt.val, result, tt.expected)

		}
	}

}
