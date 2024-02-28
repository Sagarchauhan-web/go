package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs int, rhs int) int) int {
	fmt.Printf("Running a computation with %v and %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 3 =", compute(2, 3, add))

	fmt.Println("5 - 3 =", compute(5, 3, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("Running a multiplication with %v and %v\n", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 3, mul))
}
