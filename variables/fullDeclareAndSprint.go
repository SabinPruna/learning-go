package variables

import "fmt"

var x2 = 42
var y2 = "James Bond"
var z2 = true

func assignAllToString() {
	s := fmt.Sprint(x2, y2, z2)
	fmt.Println(s)
}
