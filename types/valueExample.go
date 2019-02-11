package valueexample

import "fmt"

const (
	constA     = 1
	constB int = 2
)

const (
	year2020 = iota + 2020
	year2021
	year2022
	year2023
)

func main() {

	fmt.Println(year2020)
	fmt.Println(year2021)
	fmt.Println(year2022)
	fmt.Println(year2023)

	number := 30
	siftedNumber := number << 1
	secondNumber := 40

	fmt.Printf("%d\n", number)
	fmt.Printf("%b\n", number)
	fmt.Printf("%#x\n", number)

	fmt.Printf("%d\n", siftedNumber)
	fmt.Printf("%b\n", siftedNumber)
	fmt.Printf("%#x\n", siftedNumber)

	equals := number == secondNumber
	lessThan := number <= secondNumber
	greaterThan := number >= secondNumber
	notEqual := number != secondNumber
	lessS := number < secondNumber
	greaterS := number > secondNumber

	fmt.Println(equals, lessThan, greaterThan, notEqual, lessS, greaterS)

	stringLiteral := `"sub baby"
	
	muie vali`

	fmt.Println(stringLiteral)

}
