package sudokusolver

import "fmt"

//Solve is used to solve sudoku
func Solve(numbers ...uint16) bool {

	a := [3][9]uint16{}

	for index, value := range numbers {
		i := index / 9
		j := index % 9

		row := a[0][i]
		column := a[1][j]
		square := a[2][(3*(i/3))+j/3]
		position := value

		if (row & (1 << position)) == (1 << position) { // 100010111 000000100 -> 000000100 else 0
			fmt.Println(index, position, "row")
			return false
		} else {
			a[0][i] = row | (1 << position)
		}

		if (column & (1 << position)) == (1 << position) {
			fmt.Println(index, position, "column")
			return false
		} else {
			a[1][j] = column | (1 << position)
		}

		if (square & (1 << position)) == (1 << position) {
			fmt.Println(index, position, "square")
			return false
		} else {
			a[2][(3*(i/3))+j/3] = square | (1 << position)
		}
	}

	return true

}
