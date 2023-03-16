package main

import (
	"fmt"
	"reflect"
)

func main() {
	days := 365.2425

	var answer float64 = 42
	fmt.Println(reflect.TypeOf(days))
	fmt.Println(reflect.TypeOf(answer))
}
