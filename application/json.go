package application

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSONMarshal converts to json a slice of Users
func JSONMarshal() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	Users := []User{u1, u2, u3}

	fmt.Println(Users)

	// your code goes here
	bs, err := json.Marshal(Users)

	if err != nil {
		fmt.Println("There was an error converting to json")
	}

	fmt.Println(string(bs))
}

//Person is a struct used for unmarshalling
type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

//JSONUnmarshal unmarshells users
func JSONUnmarshal() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var persons []Person

	err := json.Unmarshal([]byte(s), &persons)

	if err != nil {
		fmt.Println("There was a problem unmarshaling the blob")
	}

	for _, person := range persons {
		fmt.Println(person)
	}

}

//EncodeUser struct used for encoding
type EncodeUser struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//JSONEncode encodes users
func JSONEncode() {

	u1 := EncodeUser{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := EncodeUser{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := EncodeUser{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	EncodeUsers := []EncodeUser{u1, u2, u3}

	fmt.Println(EncodeUsers)

	// your code goes here

	err := json.NewEncoder(os.Stdout).Encode(EncodeUsers)

	if err != nil {
		fmt.Println(err.Error())
	}
}
