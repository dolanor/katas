package fizzbuzz_test

import (
	"testing"

	"github.com/dolanor/fizzbuzz"
)

func TestFizzbuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
	}

	for _, c := range cases {
		got := fizzbuzz.Fizzbuzz(c.in)
		if got != c.want {
			t.Errorf("want %s, got %s", c.want, got)
		}
	}
}
