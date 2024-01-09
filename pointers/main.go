package main

import "fmt"

func zero(ival int) {
	ival = 0
}

func zeroPointer(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println(i)
	zero(i)
	fmt.Println(i)
	// The &i syntax gives the memory address of i
	zeroPointer(&i)
	fmt.Println(i)

	fmt.Println(&i)
}