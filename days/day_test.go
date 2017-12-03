package days

import (
	"testing"
)

func TestStringIntSlice(t *testing.T) {
	cases := []struct {
		Input  string
		Result []int
	}{
		{"123212", []int{1, 2, 3, 2, 1, 2}},
		{"777777", []int{7, 7, 7, 7, 7, 7}},
	}
	for caseNumber, c := range cases {
		r := StringToIntSlice(c.Input)
		for i, v := range r {
			if v != c.Result[i] {
				t.Errorf("%v expected, got %v in testcase: %v", c.Result, r, caseNumber)
			}
		}
	}
}
