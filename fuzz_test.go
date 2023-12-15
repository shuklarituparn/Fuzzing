package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, testcase := range testcases {
		f.Add(testcase)
	}
	f.Fuzz(func(t *testing.T, original string) { //testing as par
		rev, err1 := reverse(original)
		if err1 != nil {
			return
		}
		doubleRev, err2 := reverse(rev)
		if err2 != nil {
			return
		}
		if original != doubleRev {
			t.Errorf("Before: %q, After: %q", original, doubleRev)
		}
		if utf8.ValidString(original) && !utf8.ValidString(rev) {
			t.Errorf("Reversing produced an invalid UTF-8 string %q", rev)
		}

	})
}

//All the test pass
