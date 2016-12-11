package main

import (
	"errors"
	"fmt"
)

const (
	x = iota // x = 0
	y = iota // y = 1
	z = iota // z = 2
	w
)

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちは世界\n")

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
	m := " world"
	a := s + m
	fmt.Println(a)
	fmt.Println(x, y, z, w)
	// define a two-dimensional array with 2 elements, and each element has 4 elements
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// can b done easier
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray)
	fmt.Println(easyArray)

	// use sring as the key value, int as the type value and `make` initialize it.
	numbers := make(map[string]int)
	numbers["one"] = 1 // assign key-value
	fmt.Println(len(numbers))
	// initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	cSharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("We have the rating of C# and it is ", cSharpRating)
	} else {
		fmt.Println("We have no rating associaed with C#.")
	}

	err := errors.New("grrrrr")
	if err != nil {
		fmt.Println(err)
	}
}
