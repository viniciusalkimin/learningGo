package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//no signal positives uint8(byte), uint16(short), uint32(int), uint64(long)
	var bt byte = 23
	var ui64 uint64 = math.MaxUint64
	fmt.Println("Literal integer of", reflect.TypeOf(bt))
	fmt.Println("The max value of i64 is:", ui64)

	//with signal positives ... int8, int16, int32, int64
	i64 := math.MaxInt64
	fmt.Println("Literal integer of", reflect.TypeOf(i64))
	fmt.Println("The max value of i64 is:", i64)

	var a int
	var b float64
	var c bool
	var d string
	fmt.Printf("%v, %v, %v, %q", a, b, c, d)
}
