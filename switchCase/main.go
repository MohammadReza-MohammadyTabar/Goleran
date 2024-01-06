package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("zero")
	}
	whatAmI(true)
	whatAmI("ghrfgs")
	whatAmI(1354)
	whatAmI(341.1324)
}

func whatAmI(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("bool")
	case int :
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("%T\n", t)
	}
}
