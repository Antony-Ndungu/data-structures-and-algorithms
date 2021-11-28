package algorithms

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"MADAM", true},
		{"Madam", true},
		{"REFER", true},
		{"LONELY TYLENOL", true},
		{"ANTONY", false},
	}

	for _, c := range cases {
		got := IsPalindrome(c.in)
		if got != c.want {
			t.Errorf("isPalindrome(%q) returns %v, want %v", c.in, got, c.want)
		}
	}
}
