package main

import "fmt"

func terraform_func() {

	planets := planets{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	result := planets.terraform()
	fmt.Print(result)
}

func (p planets) terraform() planets {
	for i := range p {
		p[i] = "New " + p[i]
	}
	return p
}
