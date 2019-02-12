package groupingdata

import "fmt"

func groupingData() {
	intArray := [5]int{0, 1, 2, 3}
	intArray[4] = 4

	for index := 0; index < len(intArray); index++ {
		fmt.Println(intArray[index])
	}

	fmt.Printf("%T", intArray)

	sliceOfIntArray := []int{}

	for _, integerOfArray := range intArray {
		sliceOfIntArray = append(sliceOfIntArray, integerOfArray)
	}

	sliceOfIntArray = append(sliceOfIntArray, sliceOfIntArray...)

	customSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	customeSlice1 := customSlice[:5]
	customeSlice2 := customSlice[5:]
	customeSlice3 := customSlice[2:7]
	customeSlice4 := customSlice[1:6]

	fmt.Println(customeSlice1, customeSlice2, customeSlice3, customeSlice4)

	customSlice = append(customSlice, 52)
	customSlice = append(customSlice, 53, 54, 55)

	otherSlice := []int{56, 57, 58, 59, 60}

	customSlice = append(customSlice, otherSlice...)

	customSliceNew := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	customSliceNew = append(customSliceNew[:3], customSliceNew[6:]...)

	fmt.Println(customSliceNew)

	states := make([]string, 50, 500)
	states = append(states,
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`,
		` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`,
		` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`)

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	sliceOfSlice := [][]string{
		[]string{"test", "test2"}}
	sliceOfSlice = append(sliceOfSlice, []string{"test3", "test4"})

	for index, slice := range sliceOfSlice {
		fmt.Println(index)
		for _, element := range slice {
			fmt.Printf("%v\t", element)
		}
		fmt.Println()
	}
	preferences := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"}}

	for index := range preferences {
		fmt.Printf("At index %v the preferences are %v\n", index, preferences[index])
	}

	preferences["me"] = []string{"go"}

	for index := range preferences {
		fmt.Printf("At index %v the preferences are %v\n", index, preferences[index])
	}

	delete(preferences, "me")
}
