package functions

import (
	"fmt"
	"math"
)

//SQUARE struct
type SQUARE struct {
	length float64
}

//CIRCLE struct
type CIRCLE struct {
	radius float64
}

func (c CIRCLE) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s SQUARE) area() float64 {
	return s.length * s.length
}

//Shape interface
type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

//ReturnsInt ReturnsInt
func ReturnsInt() int {
	return 1
}

//ReturnsIntAndString ReturnsIntAndString
func ReturnsIntAndString() (int, string) {
	return 1, "hello"
}

//Foo returns sum
func Foo(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//Bar does same shit as foo
func Bar(sliceOfInts []int) int {
	var total int

	for _, number := range sliceOfInts {
		total += number
	}

	return total
}

//ILikeDefer does defering :)
func ILikeDefer() {
	defer fmt.Println("defered")

	fmt.Println("1st message")
}

func outerFunc() {
	func() {
		a := Foo
		fmt.Println(a(1, 2, 3))
	}()
}

//ReturnsFunc returns a func
func ReturnsFunc() func() int {
	return func() int {
		return 1
	}
}

func callback(function func(xi ...int) int, si ...int) int {
	return function(si...)
}
