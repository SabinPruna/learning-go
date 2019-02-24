package structs

import "fmt"

//Vehicle type
type Vehicle struct {
	Doors int
	Color string
}

//Truck type
type Truck struct {
	Vehicle
	FourWheel bool
}

//Sedan type
type Sedan struct {
	Vehicle
	Luxury bool
}

//AssignCarTypes function
func AssignCarTypes() {
	t := Truck{
		Vehicle: Vehicle{
			Doors: 2,
			Color: "white",
		},
		FourWheel: true,
	}

	s := Sedan{
		Vehicle: Vehicle{
			Doors: 4,
			Color: "black",
		},
		Luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.Doors)
	fmt.Println(s.Vehicle.Doors)
}

//AnonymousStructExample function
func AnonymousStructExample() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
