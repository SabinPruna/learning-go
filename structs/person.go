package structs

import (
	"fmt"
)

//Person struct : learning about exporting structs
type Person struct {
	First   string
	Last    string
	Flavors []string
	Age     int
}

//Speak makes the person speak
func (person Person) Speak() {
	fmt.Printf("I am %v and i am %v years old", person.First, person.Age)
}

//MakePrintPerson does what it says
func MakePrintPerson() {
	person := Person{
		First:   "JustAn",
		Last:    "Example",
		Flavors: []string{"choco", "love :)"},
	}

	fmt.Printf("This is %v %v and his favourite icecream tastes like: ", person.First, person.Last)

	for _, flavor := range person.Flavors {
		fmt.Printf("%v ", flavor)
	}
}

//MakeTwoPersons dd
func MakeTwoPersons() []Person {
	persons := []Person{}

	person1 := Person{
		First:   "John",
		Last:    "Doe",
		Flavors: []string{"something", "something else"},
	}

	persons = append(persons, person1, Person{
		First:   "JustAn",
		Last:    "Example",
		Flavors: []string{"choco", "love :)"},
	})

	return persons
}

//MapPersons functions
func MapPersons(persons []Person) map[string]Person {
	personMap := map[string]Person{}

	for _, person := range persons {
		personMap[person.Last] = person
	}

	return personMap
}

//IteratePersonMap function
func IteratePersonMap(persons map[string]Person) {
	for key, value := range persons {
		fmt.Println(persons[key])
		fmt.Println(value)
	}
}

func changeMe(person *Person) {
	*person = Person{
		Age:     1,
		First:   "changed",
		Flavors: []string{"changedFlavour"}}
}
