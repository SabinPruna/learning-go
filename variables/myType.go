package variables

import "fmt"

type myType int

var xMyType myType

func myTypeAssign() {
	fmt.Println(xMyType)
	fmt.Printf("%T\n", xMyType)
	xMyType = 42
	fmt.Println(xMyType)
}
