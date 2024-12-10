package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRuleIsValid(t *testing.T) {
	tcs := []struct {
		r     rule
		pages []int
		i     int
		want  bool
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
			got, _ := valid(tc.r, tc.pages, tc.i)

			if got != tc.want {
				t.Fail()
			}
		})
	}
}

func TestSort(t *testing.T) {
	rules, _ := parse("example1.txt")

	pages := []int{97, 13, 75, 29, 47}

	sort(rules, pages)

	want := []int{97, 75, 47, 29, 13}

	if !cmp.Equal(pages, want) {
		t.Fatal(cmp.Diff(want, pages))
	}
}
