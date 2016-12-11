package main

// An interface defines a set of methods, so if a type implements
// all the methods we say that it implements the interface.

import (
	"fmt"
	"reflect"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la la la la ... ", lyrics)
}

func (h Human) String() string {
	return "Name:" + h.name + ", Age: " + strconv.Itoa(h.age) + " years, Contact: " + h.phone
}

// Employee overloads SayHi

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, and I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

func (s Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendMoney(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

type YoungChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

func LetsTest(anything interface{}) {
	fmt.Println("Can print anything? ", anything)
}

type Element interface{}
type List []Element // list of empty interfaces

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-111-xxx"}, "Harward", 100}
	sam := Employee{Human{"Sam", 35, "555-555-xxxx"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 36, "666-222-xxxx"}, "Things Ltd.", 5000}

	//define interface i
	var i Men
	// i can store student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i can store employee
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let' use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	ii := 55
	LetsTest(ii)
	var anything string = "YES WE CAN"
	LetsTest(anything)

	Bob := Human{"Bob", 39, "0000-88852-123"}
	fmt.Println("This human is: ", Bob)

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" //tring
	list[2] = Human{"Bobby", 234, "234-26456-3342"}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an in and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is a Person and its value is: %s\n", index, value)
		}
	}

	var y float64 = 3.4
	v := reflect.ValueOf(y)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	p := reflect.ValueOf(&y)
	vv := p.Elem()
	vv.SetFloat(423.23434)
	fmt.Println(vv)
}
