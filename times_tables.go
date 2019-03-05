package main

/*
	write a function given an integer that will generate a multiplication table
	from gradeschool, id:

	input: 3
	output:

		1 2 3
		2	4	6
		3	6	9
*/

func MultiplicationTable(n int) []([]int) {
	var ret = [][]int{}
	for mult := 1; mult <= n; mult++ {
		var rowSlice = []int{}
		for row := 1; row <= n; row++ {
			rowSlice = append(rowSlice, row*mult)
		}
		ret = append(ret, rowSlice)
	}

	return ret
}
