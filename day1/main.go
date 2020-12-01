package main

import (
	"fmt"
	"github.com/olliebatch/adventofcode/day1/models"
)

func removeItem(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}


func GetNumbers (input []int) []models.NumberstoTotal {
	definedTotal := 2020

	var combinedNumbers []models.NumberstoTotal

	for i,number := range input {
		ArraywoCurrent := removeItem(input,i)

		for _,addNumber := range ArraywoCurrent {
			sum := number + addNumber
			if sum == definedTotal {
				combinedNumbers = append(combinedNumbers,models.NumberstoTotal{Number1:number,Number2:addNumber})
				// continued in case there is another duplicate number in list to get both options
				continue
			}
			// doesn't equal definedTotal so continue to next number
			continue
		}
	}
	return combinedNumbers
}

func MultiplyNumbersGenerated(totals []models.NumberstoTotal) []int {
	var totalsCombined []int

	for _,totals := range totals {
		totalsCombined = append(totalsCombined,totals.Number1* totals.Number2)
	}
	return totalsCombined
}


func main() {

	inputs := GetInput()
	combined := GetNumbers(inputs)
	fmt.Println(combined)
	totalValue := MultiplyNumbersGenerated(combined)
	fmt.Println(totalValue)
}
