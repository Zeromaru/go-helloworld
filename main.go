package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init!")
}

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) growUp() {
	p.Age = p.Age + 1
}

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(gotools.Add(10, 2))
	// fmt.Println(stringhelper.Upper("Dog"))
	// fmt.Println(stringhelper.Concat("Hello", "World"))

	// const (
	// 	myFirstName = "Pinyo"
	// 	myLastName  = "Boonsit"
	// )
	// a := "string"
	// b := 10
	// c := 10.509999
	// d := true
	// fmt.Println(myFirstName)
	// fmt.Println(myLastName)
	// fmt.Println(a, b, c, d)
	// if a == "string" {
	// 	fmt.Printf("a is %s\n", a)
	// } else {
	// 	fmt.Printf("a is %s\n", a)
	// }

	// a := [...]int{1, 2, 3, 4, 5}
	// b := []int{1, 2, 3, 4, 5}
	// b = append(b, 6)
	// fmt.Println(a)
	// fmt.Println(b)
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// for i := 0; i < len(b); i++ {
	// 	fmt.Println(b[i])
	// }
	// for key, value := range b {
	// 	fmt.Println(key, value)
	// }

	// a := 10
	// b := 20
	// switch {
	// case a > 10:
	// 	fmt.Println("a is greater than 10")
	// 	fallthrough
	// case b > 10:
	// 	fmt.Println("b is greater than 10")
	// default:
	// 	fmt.Println("What is a")
	// }

	// defer fmt.Println("World")
	// defer fmt.Println("Middle")
	// fmt.Println("Hello")

	// names := map[string]int{
	// 	"a": 20,
	// 	"b": 30,
	// 	"c": 40,
	// }
	// fmt.Println(names)
	// fmt.Println(names["a"])
	// names["b"] = 50
	// fmt.Println(names["b"])
	// for k, v := range names {
	// 	fmt.Println(k, v)
	// }

	// me := Person{
	// 	FirstName: "Pinyo",
	// 	LastName:  "Boonsit",
	// 	Age:       31,
	// }
	// me.Age = 32
	// fmt.Printf("%+v\n", me)
	// people := []Person{
	// 	Person{
	// 		FirstName: "Pinyo",
	// 		LastName:  "Boonsit",
	// 		Age:       31,
	// 	},
	// 	Person{
	// 		FirstName: "AnotherF",
	// 		LastName:  "AnotherL",
	// 		Age:       15,
	// 	},
	// }
	// for _, v := range people {
	// 	fmt.Printf("%+v\n", v)
	// }

	var people []Person
	people = append(people, Person{
		FirstName: "Pinyo",
		LastName:  "Boonsit",
		Age:       31,
	})
	people = append(people, Person{
		FirstName: "AnotherF",
		LastName:  "AnotherL",
		Age:       15,
	})
	for _, p := range people {
		p.growUp()
		fmt.Println(p.fullName())
		fmt.Println(p.Age)
	}
}
