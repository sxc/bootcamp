package main

import "fmt"

func main() {

	var (
		distance = 56000000
		speed    = 100800
	)

	var newdistance, newspeed = 56000000, 100800

	const hoursOfDays = 24

	var days = distanceToMars / speedSpaceX / hoursOfDays

	fmt.Printf("need %v days to reach Mars", days)
}
