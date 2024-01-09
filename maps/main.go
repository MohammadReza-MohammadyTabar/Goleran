package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["age"] = 12
	fmt.Println("m : ", m)

	v := m["age"]
	fmt.Println(v)

	fmt.Println(len(m))

	m["k2"] = 245
	delete(m, "age")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

}
