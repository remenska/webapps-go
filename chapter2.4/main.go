package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type Skills []string
type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	specialty string
	Skills
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}
func main() {
	var P person
	P.name = "asdasdas"
	P.age = 234234
	fmt.Println(P)

	Q := person{"Tom", 24}
	fmt.Println(Q)

	// define annonymous struct and initialize at the same time
	R := struct {
		something      string
		something_else string
	}{"sdfsdfsd", "sdfsherrwepo["}
	fmt.Println(R)

	tom := person{"Tom", 18}
	bob := person{"Bob", 24}
	older, age_difference := Older(tom, bob)
	fmt.Printf(("Of %s and %s, %s is older by %d years \n"), tom.name, bob.name, older.name, age_difference)

	mark := Student{Human: Human{"Mark", 25, 120}, specialty: "Computer Science"}
	mark.Skills = []string{"anatomy", "guitar"}
	fmt.Println(mark)
}
