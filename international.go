package main

import "fmt"

func international() {
	message := "Hola Estación Espacial Internacional"

	for i := range message {
		char := message[i]

		if char >= 'a' && char <= 'z' {
			char += 13
			if char > 'z' {
				char -= 26
			}
		}

		if char >= 'A' && char <= 'Z' {
			char += 13
			if char > 'Z' {
				char -= 26
			}
		}

		fmt.Printf("%c", char)
	}
}
