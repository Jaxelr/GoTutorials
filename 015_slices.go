package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var numbers2 []int
	printSlice(numbers2)

	if numbers2 == nil {
		fmt.Printf("slice is nil")
	}

	/* create a slice */
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers3)

	/* print the original slice */
	fmt.Println("numbers3 ==", numbers3)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers3[1:4] ==", numbers3[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers3[:3] ==", numbers3[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers3[4:] ==", numbers3[4:])

	numbers4 := make([]int, 0, 5)
	printSlice(numbers4)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number5 := numbers[:2]
	printSlice(number5)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number6 := numbers[2:5]
	printSlice(number6)

	var numbers7 []int
	printSlice(numbers7)

	/* append allows nil slice */
	numbers7 = append(numbers7, 0)
	printSlice(numbers7)

	/* add one element to slice*/
	numbers7 = append(numbers7, 1)
	printSlice(numbers7)

	/* add more than one element at a time*/
	numbers7 = append(numbers7, 2, 3, 4)
	printSlice(numbers7)

	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers8 := make([]int, len(numbers7), (cap(numbers7))*2)

	/* copy content of numbers to numbers8 */
	copy(numbers8, numbers7)
	printSlice(numbers8)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
