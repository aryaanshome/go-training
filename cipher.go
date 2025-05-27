package main

import (
	"fmt"
	"strings"
)

func cipher() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	cipherText := ""

	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	for i := 0; i < len(plainText); i++ {
		plainText_char := plainText[i]

		j := i % len(keyword)

		keyword_char := keyword[j]

		message_num := plainText_char - 'A'
		keyword_num := keyword_char - 'A'

		cipher_num := (message_num+keyword_num)%26 + 'A'
		cipherText += string(cipher_num)
	}
	fmt.Println(cipherText)
}
