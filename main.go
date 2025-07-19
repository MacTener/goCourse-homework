package main

import "fmt"

const usdToEur float64 = 0.86
const usdToRub float64 = 78.52
const eurToRub float64 = usdToRub / usdToEur

func main() {
	userAnswer := userImput()
	fmt.Println("Вы ввели:", userAnswer)
}

func userImput() string {

	var answer string

	fmt.Println("Введите данные:")
	fmt.Scan(&answer)

	return answer
}

func calculate(num int, num2 float64, num3 float64) {

}
