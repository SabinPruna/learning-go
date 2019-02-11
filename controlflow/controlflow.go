package controlflow

import (
	"fmt"
	"time"
)

func controlflow() {
	for counter := 1; counter <= 10000; counter++ {
		fmt.Println(counter)
	}

	for counter := 65; counter <= 90; counter++ {
		fmt.Println(counter)
		for index := 0; index < 3; index++ {
			fmt.Printf("\t\t%U\t%q\n", counter, counter)
		}
	}

	for {
		duration := time.Since(time.Date(1998, time.April, 2, 0, 0, 0, 0, time.UTC))
		fmt.Println(duration.Seconds() / 2600640)
		break
	}

	year := 1998
	for year <= 2019 {
		fmt.Println(year)
		year++
	}

	year = 1998

	for {
		if year == 2017 {
			break
		} else {
			fmt.Println(year)
			year++
		}
	}

	for index := 10; index < 100; index++ {
		fmt.Printf("Remainder of %v devided by 4 if %v\n", index, index%4)
	}

	switch {
	case false:
		fmt.Println("won't run")
	default:
		fmt.Println("default")

	}

	favSport := "none"

	switch favSport {
	case "football":
		fmt.Println("football")
	case "tennis":
		fmt.Println("tennis")
	case "none":
		fmt.Println("none")
		fallthrough
	default:
		fmt.Println("default")

	}
}
