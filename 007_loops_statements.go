package main

import "fmt"

func main() {
	for_loop()
	nested_loop()
	break_example()
	continue_example()
	goto_example()
}

func for_loop() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}

func nested_loop() {
	/* local variable definition */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // if factor found, not prime
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

func break_example() {
	/* local variable definition */
	var a int = 10

	/* for loop execution */
	for a < 20 {
		fmt.Printf("break: value of a: %d\n", a)
		a++
		if a > 15 {
			/* terminate the loop using break statement */
			break
		}
	}
}

func continue_example() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("continue: value of a: %d\n", a)
		a++
	}
}

func goto_example() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
LOOP:
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("goto: value of a: %d\n", a)
		a++
	}
}
