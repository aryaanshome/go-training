package main

import "fmt"

func row_tables(celsius_temp celsius) farenheit {

	farenheit_temp := (celsius_temp * 9.0 / 5.0) + 32.0
	return farenheit(farenheit_temp)

}

func drawtables(r rowtables) {

	fmt.Println("========================")
	fmt.Println("|    °C     |    °F     |")
	fmt.Println("========================")
	for i := -40; i < 100; i += 5 {
		farenheit_temp := r(celsius(i))
		fmt.Printf("|   %v   |   %v   |\n", i, farenheit_temp)
	}
	fmt.Println("========================")
}
