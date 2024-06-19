package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	fmt.Println("Addition =", a+b)
	fmt.Println("Subtraction =", a-b)
	fmt.Println("Division =", a/b)
	fmt.Println("Multiplication =", a*b)
	fmt.Println("Rest of Division =", a%b)

	c := 2.0
	d := 4.0
	//operation using Math..
	fmt.Println("Max =", math.Max(float64(a), float64(b)))
	fmt.Println("Min =", math.Min(float64(a), float64(b)))
	fmt.Println("Pow =", math.Pow(c, d))

	//bitwise
	fmt.Println("E =", a&b)
	fmt.Println("Ou =", a|b)
	fmt.Println("Xor =", a^b)

}
