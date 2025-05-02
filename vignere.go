package main

import (
	"fmt"
)

func vignere() {
	var message string = "CSOITEUIWUIZNSROCNKFD"
	var keyword string = "GOLANG"
	cipher := ""

	for i := 0; i < len(message); i++ {
		message_char := message[i]

		j := i % len(keyword)

		keyword_char := keyword[j]

		message_num := message_char - 'A'
		keyword_num := keyword_char - 'A'

		cipher_num := (message_num-keyword_num+26)%26 + 'A'
		cipher += string(cipher_num)
	}
	fmt.Println(cipher)
}
