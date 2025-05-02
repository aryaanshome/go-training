package main

import (
	"fmt"
	"math/rand"
)

func int_piggy() {
	const nickel = 5
	const dimes = 10
	const quarters = 25
	var piggy_bank_limit = 2000
	var piggy_bank_amount = 0

	for piggy_bank_limit > 0 {

		var amount_of_nickel = rand.Intn(10)
		var amount_of_dimes = rand.Intn(5)
		var amount_of_quarters = rand.Intn(4)

		piggy_bank_amount += (nickel * amount_of_nickel) + (dimes * amount_of_dimes) + (quarters * amount_of_quarters)
		fmt.Printf("%v is in the Piggy bank \n", piggy_bank_amount)
		piggy_bank_limit -= piggy_bank_amount
		fmt.Printf("%v is the limit remaining Piggy bank \n", piggy_bank_limit)
	}
}
