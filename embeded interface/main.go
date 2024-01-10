package main

import "fmt"

type base struct {
	num int
}

func (b base) desc() string {
	return fmt.Sprintf("num :%v ",b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 12,
		},
		str: "salam",
	}

	fmt.Println(co.num)

	type describer interface{
		desc() string
	}
	
	var d describer=co
	fmt.Printf(d.desc())
}