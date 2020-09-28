package main

import (
	"fmt"
	"math"
)

/* define a circle */
type Circle struct {
	x, y, radius float64
}

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200

	/* calling a function to get max value */
	ret := max(a, b)

	fmt.Printf("Max value is : %d\n", ret)

	c, d := swap("Jaxel", "Rojas")
	fmt.Println(c, d)

	fmt.Println("Before Swapping: ", a, b)
	swap_by_reference(&a, &b)
	fmt.Println("After Swapping: ", a, b)

	/* declare a function variable */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* use the function */
	fmt.Println(getSquareRoot(9))

	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* return multiple values */
func swap(x, y string) (string, string) {
	return y, x
}

func swap_by_reference(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}

/* function closure  */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
