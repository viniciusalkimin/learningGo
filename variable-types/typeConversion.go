package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Prints AskTable reference value for (125)
	fmt.Println("Test" + string(125))

	//Int to String
	fmt.Println("Test" + strconv.Itoa(123))

	//String to Int or Error
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("is True!")
	} else {
		fmt.Println("is False!")
	}
}
