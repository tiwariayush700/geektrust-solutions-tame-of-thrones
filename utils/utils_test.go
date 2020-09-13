package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateAntiClockwise(t *testing.T) {
	tests := []struct {
		val      string
		d        int32
		expected string
	}{
		{
			val:      "ABCDEF",
			d:        5,
			expected: "vwxyza",
		},
		{
			val:      "WXYZA",
			d:        5,
			expected: "rstuv",
		},
		{
			val:      "TESTIng",
			d:        0,
			expected: "testing",
		},
		{
			val:      "ABCDE",
			d:        27,
			expected: "zabcd",
		},
		{
			val:      "ABCDE",
			d:        26,
			expected: "abcde",
		},
		{
			val:      "abc",
			d:        28,
			expected: "yza",
		},
	}
	for _, test := range tests {
		actual := RotateAntiClockwise(test.val, test.d)
		assert.Equal(t, test.expected, actual)
	}
}
