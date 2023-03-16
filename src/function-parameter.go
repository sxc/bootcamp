package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// A function as a parameter
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vºK\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	measureTemperature(3, fakeSensor)
}
