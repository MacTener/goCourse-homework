package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("___КАЛЬКУЛЯТОР___")

	for {

		userInput, err := getOperation()
		if err != nil {
			fmt.Println("Несуществующая операция. Повторите ввод")
			continue
		}

		userNums := getUserNums()

		sliceNumbers, err := stripString(userNums)
		if err != nil {
			fmt.Println("Ошибка введеных данных. Повторите ввод")
			continue
		}

		answer, err := calculate(sliceNumbers, userInput)
		if err != nil {
			fmt.Println("Ошибка введеных данных. Повторите ввод")
			continue
		}

		fmt.Println("Ответ:", answer)

		break
	}

}

func getOperation() (string, error) {

	fmt.Println("Введите операцию:")
	fmt.Println("Доступные операции - AVG, SUM, MED")

	var userInput string
	fmt.Scan(&userInput)

	if userInput != "AVG" && userInput != "SUM" && userInput != "MED" {
		return "", errors.New("INPUT_ERROR")
	}
	return userInput, nil

}

func getUserNums() string {

	fmt.Println("Введите числа: (прим. 1,2,3,4)")

	var userInput string
	fmt.Scan(&userInput)

	return userInput

}

func stripString(numString string) ([]int, error) {

	var sliceNumbers []int

	numArray := strings.Split(numString, ",")
	for _, value := range numArray {

		value, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Неккоректные данные")
			return nil, errors.New("STRING_ARRAY_ERROR")
		}

		sliceNumbers = append(sliceNumbers, int(value))

	}

	return sliceNumbers, nil

}

func calculate(numArray []int, operation string) (int, error) {

	switch operation {
	case "AVG":

		var numCount int
		var numSumm int

		for _, value := range numArray {
			numSumm += value
			numCount += 1
		}

		average := numSumm / numCount
		return average, nil

	case "SUM":

		var numSumm int

		for _, value := range numArray {
			numSumm += value
		}
		return numSumm, nil

	case "MED":

		sort.Ints(numArray)
		lenArray := len(numArray)

		if lenArray%2 == 0 {
			medNum := numArray[lenArray/2]
			return medNum, nil
		} else {
			middleNum := lenArray / 2
			midArifm := (numArray[middleNum] + numArray[middleNum-1]) / 2
			return midArifm, nil
		}

	default:
		return 0, errors.New("MISSING_ERROR")
	}

}
