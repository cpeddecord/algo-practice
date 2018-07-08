package main

import (
	"testing"
)

type Assert struct {
	A        string
	B        string
	C        string
	Expected bool
}

func TestInterleaved(t *testing.T) {
	asserts := []Assert{
		{
			"abc",
			"def",
			"abcdef",
			true,
		},
		// {
		// 	"abc",
		// 	"def",
		// 	"defabc",
		// 	true,
		// },
		// {
		// 	"ab",
		// 	"ac",
		// 	"acba",
		// 	false,
		// },
	}

	for _, a := range asserts {
		result := Interleave(a.A, a.B, a.C)
		if result != a.Expected {
			t.Error(a.A, a.B, a.C)
		}
	}

}
