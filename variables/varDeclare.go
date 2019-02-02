package variables

import "fmt"

var x int
var y string
var z bool

func varDeclare() {
	//prints the ZERO VALUE of the variables declared
	fmt.Println(x, y, z)
}
