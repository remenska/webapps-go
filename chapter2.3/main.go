package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (add int, multiplied int) {
	add = A + B
	multiplied = A * B
	return
}

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		if index == 5 {
			continue
		}
		sum += index
	}

	fmt.Println("Sum is equal to", sum)
	integer := 6
	switch integer {
	case 4:
		fmt.Println("integer <= 4")
		fallthrough
	case 5:
		fmt.Println("integer <= 5")
		fallthrough
	case 6:
		fmt.Println("integer <= 6")
		fallthrough
	case 7:
		fmt.Println("integer <= 7")
		fallthrough
	case 8:
		fmt.Println("integer <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_xz := max(x, z)
	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
	fmt.Println(SumAndProduct(x, y))
	myfunc(1, 2, 3, 4, 5)

	_ = add1(&x)
	fmt.Println("x = ", x)
	// fmt.Println("x1 = ", x1)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	even := filter(slice, isEven)
	fmt.Println("Odd elements of slice are: ", odd)
	fmt.Println("Even elements of slice are: ", even)
}

func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func add1(a *int) int {
	*a = *a + 1
	return *a
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
