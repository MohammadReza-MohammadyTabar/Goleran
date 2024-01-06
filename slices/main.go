package main

import (
	"fmt"
	"slices"
)

func main() {
	// start of array
	// array should contain number of element and type of them
	var a [5]int
	fmt.Println(a)

	a[4] = 10
	fmt.Println(a)
	fmt.Println(a[4])

	b := [3]int{1, 3, 5}
	fmt.Println(b)

	for key, val := range b {
		fmt.Println(key, val)
	}

	// if don't want to use key replace it with _
	for _, val := range a {
		fmt.Println(val)
	}
	// end of array

	// start slices
	// slice only have type of element and not number of them
	fmt.Println("*********************slices*******************")
	var c []string
	fmt.Println("uninit:", c, c == nil, len(c) == 0)

	c=make ([]string , 2)
	fmt.Println(c,len(c),cap(c))

	c[0]="mohammad"
	c=append(c, "ali")
	c=append(c, "reza","mahmud")
	fmt.Println(c)
	fmt.Println(c,len(c),cap(c))
	
	s:=make([]string , len(c))
	copy(s,c)
	fmt.Println("s : ",s)

	c=append(c, "javad")
	fmt.Println(c)
	fmt.Println(s)
	
	l:=s[2:5]
	fmt.Println(l)
	l=s[:3]
	fmt.Println(l)
	l=s[2:]
	fmt.Println(l)
	l=s[0:]

	if slices.Equal(s,l ) {
        fmt.Println("s == l")
    }

	if slices.Equal(s,c ) {
        fmt.Println("s == c")
    }


	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }

	fmt.Println("2d: ", twoD)
}
