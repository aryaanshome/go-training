package main

import (
	"fmt"
)

func caesar() {
	quote := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(quote); i++ {
		char := quote[i]

		if char >= 'a' && char <= 'z' {
			char -= 3
			if char < 'a' {
				char += 26
			}
		}

		if char >= 'A' && char <= 'Z' {
			char -= 3
			if char < 'A' {
				char += 26
			}
		}

		fmt.Printf("%c", char)
	}
}
