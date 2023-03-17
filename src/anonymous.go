package main

import "fmt"

func main() {
	func() { // Declares an anonymous function
		fmt.Println("Functions anonymous")
	}() // Invokes the function
}
