package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3, 2, 1}
	// results := DuplicateArr(nums, 3)
	// fmt.Println(results)

	// nums2 := []int{1, 2, 3, 4, 1, 5}
	// results2 := DuplicateArr(nums2, 4)
	// fmt.Println(results2)

	// table := MultiplicationTable(12)
	// fmt.Println(table)
	
	html := StyleParse("Hello,world", []Tag{
		{0, 2, "i"},
		{4, 8, "b"},
		{7, 10, "u"},
	})

	fmt.Println(html)
}
