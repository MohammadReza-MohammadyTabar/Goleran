package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func sumV2(x int, y int) {
	fmt.Println(x + y)
}
func main() {

	fmt.Println(sum(2, 2))
	sumV2(2, 3)

}
