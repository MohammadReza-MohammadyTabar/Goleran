package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

//receiver by refrence
func (p *person) printPerson(){
	fmt.Println("name :",&p.name,"\nage :",&p.age)
	fmt.Println("name :",p.name,"\nage :",p.age)
}

//receiver by value
func (p person) printPerson2(){
	// using pass by value make it copy of data so the address is changed
	fmt.Println("name :",&p.name,"\nage :",&p.age)
}
func main() {
	fmt.Println(newPerson("mohammad",12))
	fmt.Println(&person{name: "Ann", age: 40})

	s:=person{"ali",45}
	//using reciver to print perosn 
	fmt.Println("name :",&s.name,"\nage :",&s.age)
	s.printPerson()
	s.printPerson2()
	sp:=&s
	fmt.Println(sp.name)
	dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}