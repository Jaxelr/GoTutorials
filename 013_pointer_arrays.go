package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var aptr [MAX]*int

	for i = 0; i < MAX; i++ {
		aptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *aptr[i])
	}

	/* pointer to pointer */
	var ba int
	var ptr *int
	var pptr **int

	ba = 3000

	/* take the address of var */
	ptr = &ba

	/* take the address of ptr using address of operator & */
	pptr = &ptr

	/* take the value using pptr */
	fmt.Printf("Value of ba = %d\n", ba)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)

	/* local variable definition */
	var c int = 100
	var d int = 200

	fmt.Printf("Before swap, value of c : %d\n", c)
	fmt.Printf("Before swap, value of d : %d\n", d)

	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap(&c, &d)

	fmt.Printf("After swap, value of c : %d\n", c)
	fmt.Printf("After swap, value of d : %d\n", d)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}
