package main

import (
	"fmt"
	"math/rand"
)

func piggy() {
	const nickel float64 = 0.05
	const dimes float64 = 0.10
	const quarters float64 = 0.25
	var piggy_bank_limit float64 = 20
	var piggy_bank_amount float64 = 0

	for piggy_bank_limit > 0 {

		var amount_of_nickel = rand.Intn(10) + 1
		var amount_of_dimes = rand.Intn(5) + 1
		var amount_of_quarters = rand.Intn(4) + 1

		piggy_bank_amount += (nickel * float64(amount_of_nickel)) + (dimes * float64(amount_of_dimes)) + (quarters * float64(amount_of_quarters))
		fmt.Printf("%.2f is in the Piggy bank \n", piggy_bank_amount)
		piggy_bank_limit -= piggy_bank_amount
		fmt.Printf("%.2f is the limit remaining Piggy bank \n", piggy_bank_limit)
	}
}
