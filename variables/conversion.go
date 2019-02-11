package variables

import "fmt"

type myConversionType int

var xConv myType
var yConv int

func conversionOfMyType() {
	fmt.Println(x)
	fmt.Printf("%T\n", xConv)
	xConv = 42
	fmt.Println(xConv)
	yConv = int(xConv)
	fmt.Println(y)
	fmt.Printf("%T\n", yConv)
}


