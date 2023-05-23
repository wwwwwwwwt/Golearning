package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("None")
	}

	WhoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am bool!")
		case int:
			fmt.Println("i am int!")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	WhoAmI(true)
	WhoAmI(10)
	WhoAmI("hey")
}
