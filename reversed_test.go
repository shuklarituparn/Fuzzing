package main

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
		{"!12345", "54321!"},
	}
	for _, tc := range testCases {
		rev := reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want: %q", rev, tc.want)
		}
	}
}

//Now to convert unit test to fuzztest
