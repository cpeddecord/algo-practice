package main

import (
	"testing"
)

type DuplicateAssert struct {
	arr      []int
	k        int
	Expected bool
}

func TestDuplicateArr(t *testing.T) {
	asserts := []DuplicateAssert{
		{
			[]int{1, 2, 3, 2, 1},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 4, 1, 5},
			4,
			false,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1, 2},
			2,
			false,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1, 2},
			3,
			true,
		},
	}

	for _, a := range asserts {
		result := DuplicateArr(a.arr, a.k)
		if result != a.Expected {
			t.Error(a.arr, a.k)
		}
	}

}
