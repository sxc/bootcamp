package main

import "fmt"

func main() {
	greeting("Jim")
	g := greeting
	fmt.Println("Function type: %T\n", g)
	g("Jim")
	greetingMultiple("Jim", "Bob", "Tim")

}

func greeting(subject string) {
	fmt.Println("Function hello", subject)
}

func greetingMultiple(subjects ...string) {
	for _, subject_ := range subjects {
		fmt.Println("Function hello", subject_)
	}
}
