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

		userNumber, err := startNumber()
		if err != nil {
			fmt.Println("Введено неверное число!")
			continue
		}

		userEndAnswer, err := endCurrency(userAnswer)
		if err != nil {
			fmt.Println("Введена невереная валюта!")
			continue
		}

		convertedValue := calculate(userNumber, userAnswer, userEndAnswer)
		fmt.Println("Полученное значение:", convertedValue)
		fmt.Println("Хотите продолжить? (y/n)")

		var userContinue string
		fmt.Scan(&userContinue)

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
	} else if answer == userCurrency {
		return "", errors.New("SAME_CURRENCY")
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
