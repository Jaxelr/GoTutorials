package main

import "fmt"

func main() {
   var x float64 = 20.0

   y := 42 //Notice this declaration for type inference
   fmt.Println(x)
   fmt.Println(y)
   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)	
}