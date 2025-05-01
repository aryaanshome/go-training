package main

import (
	"fmt"
)

func printlightspeed() {
	const lightspeed = 27792 //km
	var distance = 56000000  //km

	fmt.Printf("%v seconds \n", distance/lightspeed)

	distance = 41000000 //km
	fmt.Printf("%v seconds", distance/lightspeed)
}
