package main

import "fmt"

const usdToEur float64 = 0.86
const usdToRub float64 = 78.52
const eurToRub float64 = usdToRub / usdToEur

func main() {

	var startEur float64 = 50.0
	rub := startEur * eurToRub

	fmt.Print(rub)

}
