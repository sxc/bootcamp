package main

import "fmt"

func main() {
	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "second")

	const distanceToMars = 96300000

	var speedSpaceX = 100800

	const hoursOfDays = 24

	var days = distanceToMars / speedSpaceX / hoursOfDays

	fmt.Printf("need %v days to reach Mars", days)
}
