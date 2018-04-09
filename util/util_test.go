package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	cases := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{1, 2},
			expect: 1,
		},
		{
			input:  []int{2, 1},
			expect: 1,
		},
		{
			input:  []int{2, 2},
			expect: 2,
		},
	}
	for _, c := range cases {
		val := Min(c.input[0], c.input[1])
		assert.Equal(t, c.expect, val)
	}
}
