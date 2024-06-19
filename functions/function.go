package main

import "fmt"

func sum(a int64, b int64) int64 {
	return a + b
}

func print(value int64) {
	fmt.Println(value)
}

func main() {
	result := sum(2, 4)
	print(int64(result))
}
