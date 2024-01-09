package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func sumV2(x int, y int) {
	fmt.Println(x + y)
}

func withMulipleReturnVal() (int, string) {
	return 12, "mohammad"
}

// Variadic function
// ...int => []int
func v(nums ...int) {
	tot := 0
	for _, num := range nums {
		tot += num
	}
	fmt.Println(tot)
}

// Closures =>  anonymous functions
// Anonymous functions are useful when you want to define a function inline without having to name it.
func intSec() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}


// recursive function 
func re(x int )int {
	if x ==0{
		return 1
	}
	return x + re(x-1)
}
func main() {

	fmt.Println(sum(2, 2))
	sumV2(2, 3)

	age, name := withMulipleReturnVal()
	fmt.Println("age :", age, "\nname :", name)
	fmt.Println(withMulipleReturnVal())

	// Variadic function
	v(1)
	v(1, 2)
	v(3, 4, 5, 6, 7)

	// Closures =>  anonymous functions
	nextInt := intSec()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSec()
	fmt.Println(newInts())
	//recursive 
	fmt.Println(re(2))
}
