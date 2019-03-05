package application

import (
	"fmt"
	"sort"
)

//SortIntsAndStrings sorts ints and string
func SortIntsAndStrings() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)

	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

//User is struct
type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//ByAge is used for sorting users by age
type ByAge []User

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

//ByLast is used for sorting users by last
type ByLast []User

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

//SortUsers sorts users
func SortUsers() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	for _, user := range users {
		sort.Strings(user.Sayings)
	}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)

}
