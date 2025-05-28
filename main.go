package main

type celsius float64
type kelvin float64
type farenheit float64
type planets []string
type array [8][8]string
type wordcounter map[string]int

type rowtables func(celsius) farenheit
type revrowtables func(farenheit) celsius

func main() {
	// fmt.Println("Hello World")
	// printlightspeed()
	//random()
	//iftest()
	//guess()
	//float_piggy()
	//caesar()
	//international()
	// vignere()
	//temperature()
	// sensor := calibrate(realSensor, 5)
	// fmt.Println(sensor())
	//cipher()
	// drawtables(row_tables)
	// drawrevtables(rev_row_tables)
	// chess()
	// terraform_func()
	wordsextracter()
}
