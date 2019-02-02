package variables

import (
	"fmt"
)

func variableAssign() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)

	fmt.Printf("%T\t%T\t%T\n", x, y, z)
	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Printf("%#v\t%#v\t%#v\n", x, y, z)
}
