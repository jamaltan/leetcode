package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		param  string
		result int
	}{
		{
			param:  "abcabcbb",
			result: 3,
		},
		{
			param:  "bbbb",
			result: 1,
		},
		{
			param:  "pwwkew",
			result: 3,
		},
		{
			param:  "abba",
			result: 2,
		},
		{
			param:  "au",
			result: 2,
		},
		{
			param:  "aab",
			result: 2,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, lengthOfLongestSubstring(tt.param), tt.param+" should be equal")
	}
}
