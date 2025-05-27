package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type farenheit float64

func temperature() {

	kelvin_temp := kelvin(0.0)
	celsius_temp := kelvintocelsius(kelvin_temp)

	fmt.Printf("%v C is the temperature in Celsius \n", celsius_temp)

	farenheit_temp := celsiustofarenheit(celsius(celsius_temp))
	fmt.Printf("%6.2f F is the temperature in Farenheit", farenheit_temp)

}
