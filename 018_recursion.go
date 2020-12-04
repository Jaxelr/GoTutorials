package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	var i int = 15
	fmt.Printf("Factorial of: %d is %d \n", i, factorial(i))

	fmt.Printf("Fiboaci: ")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}
