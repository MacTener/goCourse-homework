package main

import "fmt"

const usdToEur float64 = 0.86
const usdToRub float64 = 78.52

func main() {

	var startEur float64 = 50.0
	usd := startEur / usdToEur
	rub := usd * usdToRub

	fmt.Print(rub)

}
