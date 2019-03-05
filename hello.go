package main

import (
	"github.com/sabinpruna/application"
)

func main() {
	//structs.AssignCarTypes()

	//solve sudoku using bitwise operations
	//for each line allocate a number in base 2

	/*
		a[0] -> eight rows , 0 or 1 for each line if value was present (array of bites)
		a[1] -> cols
		 -> a[2] squares
	*/

	// k = i/3 * 3 + j/3 find which square you're in
	//for 1d slice  t => i = t/9, j = t%9
	// x is 8 bits from 0 to 8

	/**
	 * check if poz of x is 1 or 0, i'd have 1000000, where 0 is poz, 2^poz
	 * x & (1 << poz) == 0 it's good else it's bad and x = x | (1 << poz)
	 *

	**/
	// numbers := []uint16{
	// 	5, 3, 4, 6, 7, 8, 9, 1, 2,
	// 	6, 7, 2, 1, 9, 5, 3, 4, 8,
	// 	1, 9, 8, 3, 4, 2, 5, 6, 7,
	// 	8, 5, 9, 7, 6, 1, 4, 2, 3,
	// 	4, 2, 6, 8, 5, 3, 7, 9, 1,
	// 	7, 1, 3, 9, 2, 4, 8, 5, 6,
	// 	9, 6, 1, 5, 3, 7, 2, 8, 4,
	// 	2, 8, 7, 4, 1, 9, 6, 3, 5,
	// 	3, 4, 5, 2, 8, 6, 1, 7, 9}

	//fmt.Println(sudokusolver.Solve(numbers...))

	//structs
	application.SortUsers()

}
