package main

import "testing"

func TestRuleIsValid(t *testing.T) {
	tcs := []struct {
		r        rule
		pages    []int
		i        int
		expected bool
	}{
		{
			rule{47, 53},
			[]int{75, 47, 61, 53, 29},
			1,
			true,
		},
		{
			rule{47, 53},
			[]int{75, 47, 61, 53, 29},
			3,
			true,
		},
		{
			rule{97, 75},
			[]int{75, 97, 47, 61, 53},
			0,
			false,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			if valid(tc.r, tc.pages, tc.i) != tc.expected {
				t.Fail()
			}
		})
	}
}
