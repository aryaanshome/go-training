package main

import "fmt"

func iftest() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("Go up the mountain")
	} else if command == "go west" {
		fmt.Println(("Go down the mountain"))
	} else {
		fmt.Println("Jump off the mountain")
	}
}
