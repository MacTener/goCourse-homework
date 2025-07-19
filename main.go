package main

import (
	"errors"
	"fmt"
)

const usdToEur float64 = 0.86
const usdToRub float64 = 78.52
const eurToRub float64 = usdToRub / usdToEur

func main() {

	fmt.Println("__Конвентор валют__")

	for {

		userAnswer, err := startCurrency()
		if err != nil {
			fmt.Println("Введена невереная валюта!")
			continue
		}

		// fmt.Println("=======НАЧАЛЬНАЯ ВАЛЮТА:")
		// fmt.Println("userAnswer:", userAnswer)

		userNumber, err := startNumber()
		if err != nil {
			fmt.Println("Введено неверное число!")
			continue
		}

		// fmt.Println("=======ЧИСЛО ДЛЯ КОНВЕНТАРЦИИ:")
		// fmt.Println("userNumber:", userNumber)

		userEndAnswer, err := endCurrency(userAnswer)

		// fmt.Println("=======НАЧАЛЬНАЯ ВАЛЮТА:")
		// fmt.Println("userEndAnswer:", userEndAnswer)

		convertedValue := calculate(userNumber, userAnswer, userEndAnswer)
		fmt.Println("Полученное значение:", convertedValue)
		fmt.Println("Хотите продолжить? (y/n)")

		var userContinue string
		fmt.Scan(&userContinue)

		// fmt.Println("------ОТВЕТ ПОЛЬЗОВАТЕЛЯ")
		// fmt.Println("userContinue:", userContinue)

		if userContinue == "y" || userContinue == "Y" {
			continue
		}

		fmt.Println("До свидания!")
		break

	}

}

func startCurrency() (string, error) {

	var answer string

	fmt.Println("Введите начальную валюту:")
	fmt.Println("Доступные валюты - USD, EUR, RUB")
	fmt.Scan(&answer)

	if answer != "USD" && answer != "EUR" && answer != "RUB" {
		return "", errors.New("CURRENCY_MISMATCH")
	}

	return answer, nil
}

func endCurrency(userCurrency string) (string, error) {

	var answer string

	fmt.Println("Введите валюту, в которую нужно конвентировать:")
	switch userCurrency {
	case "USD":
		fmt.Println("Доступные валюты для конвентариции: EUR, RUB")
	case "EUR":
		fmt.Println("Доступные валюты для конвентариции: USD, RUB")
	case "RUB":
		fmt.Println("Доступные валюты для конвентариции: EUR, USD")
	default:
		return "", errors.New("CURRENCY_MISMATCH")
	}

	fmt.Scan(&answer)
	if answer != "USD" && answer != "EUR" && answer != "RUB" {
		return "", errors.New("CURRENCY_MISMATCH")
	}

	return answer, nil
}

func startNumber() (float64, error) {

	var number float64

	fmt.Println("Введите число для конвентарции:")
	fmt.Scan(&number)

	if number < 0 {
		return 0, errors.New("VALUE_ERROR")
	}
	return number, nil

}

func calculate(num float64, startCurrency string, endCurrency string) float64 {

	// const usdToEur float64 = 0.86
	// const usdToRub float64 = 78.52
	// const eurToRub float64 = usdToRub / usdToEur

	// -> Mult
	// <- div

	var endValue float64

	switch startCurrency {
	case "USD":
		if endCurrency == "EUR" {
			endValue = num * usdToEur
		} else {
			endValue = num * usdToRub
		}

	case "EUR":
		if endCurrency == "USD" {
			endValue = num / usdToEur
		} else {
			endValue = num * eurToRub
		}

	case "RUB":
		if endCurrency == "USD" {
			endValue = num / usdToRub
		} else {
			endValue = num / eurToRub
		}
	}

	return endValue

}
