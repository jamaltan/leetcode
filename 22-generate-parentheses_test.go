package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		param  int
		result []string
	}{
		{
			param:  1,
			result: []string{"()"},
		},
		{
			param:  3,
			result: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, generateParenthesis(tt.param), "should be equal")
	}
}
