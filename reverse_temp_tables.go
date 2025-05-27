package main

import "fmt"

type revrowtables func(farenheit) celsius

func rev_row_tables(farenheit_temp farenheit) celsius {

	celsius_temp := (farenheit_temp * 5.0 / 9.0) - 32.0
	return celsius(celsius_temp)

}

func drawrevtables(r revrowtables) {

	fmt.Println("========================")
	fmt.Println("|    °F     |    °C     |")
	fmt.Println("========================")
	for i := -40; i < 100; i += 5 {
		celsius_temp := r(farenheit(i))
		fmt.Printf("|   %v   |   %4.2f   |\n", i, celsius_temp)
	}
	fmt.Println("========================")
}
