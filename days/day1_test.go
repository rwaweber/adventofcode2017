package days

import (
	"testing"
)

func TestCaptcha(t *testing.T) {
	cases := []struct {
		Problem string
		Result  int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
		{"8812228", 20},
	}

	for _, c := range cases {
		dummyd1 := Day1{
			Input: c.Problem,
		}
		r := dummyd1.Captcha()
		if r != c.Result {
			t.Errorf("%v expected, got %v", c.Result, r)
		}
	}
}
