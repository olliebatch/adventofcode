package main

import (
	"fmt"
	"github.com/olliebatch/adventofcode/day1/models"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ThreeNumber(ArraywoCurrent []int,numberLookingFor int, numb int, currentNumber int) []int {
	var numberArr []int
	updatedNumberSearch := numberLookingFor - currentNumber
	for _,thirdNumber := range ArraywoCurrent {
		if thirdNumber == updatedNumberSearch {
			numberArr = append(numberArr,numb,currentNumber,thirdNumber)
			// continued in case there is another duplicate number in list to get both options
			continue
		}
		continue
	}
	return numberArr
}


func LoopAndFindNumbers (ArraywoCurrent []int, numb int, definedTotal int, combinedNumbers []models.NumberstoTotal) ([]models.NumberstoTotal , []int) {
	numberLookingFor := definedTotal - numb
	combinedThreeNumbers := []int{}
	for _,addNumber := range ArraywoCurrent {
		if addNumber == numberLookingFor {
			combinedNumbers = append(combinedNumbers,models.NumberstoTotal{Number1:numb,Number2:addNumber})
			// continued in case there is another duplicate number in list to get both options
			continue
		}
		if addNumber < numberLookingFor {
			three := ThreeNumber(ArraywoCurrent,numberLookingFor,numb,addNumber)
			if len(three) > 0 {
				combinedThreeNumbers = append(combinedThreeNumbers,three...)
				continue
			}
			continue
		}
		// doesn't equal definedTotal so continue to next number
		continue
	}
	return combinedNumbers,combinedThreeNumbers
}


func GetNumbers (input []int) ([]models.NumberstoTotal,[]int) {
	definedTotal := 2020

	var combinedNumbers []models.NumberstoTotal
	var combinedThree []int

	for _,number := range input {

		combined, threeNumb := LoopAndFindNumbers(input,number,definedTotal,combinedNumbers)
		combinedNumbers = combined
		if len(threeNumb) > 0 {
			combinedThree = threeNumb
		}
	}
	return combinedNumbers,combinedThree
}

func MultiplyNumbersGenerated(totals []models.NumberstoTotal) []int {
	var totalsCombined []int

	for _,totals := range totals {
		totalsCombined = append(totalsCombined,totals.Number1* totals.Number2)
	}
	return totalsCombined
}


func MultiplyArray(totals []int) int {
	var sum int

	for i,value := range totals {
		if i == 0 {
			sum = value
			continue
		}
		sum = sum * value
	}
	return sum
}


func main() {

	inputs := GetInput()
	combined,three := GetNumbers(inputs)
	threeUnique := unique(three)
	sumThree := MultiplyArray(threeUnique)
	totalValue := MultiplyNumbersGenerated(combined)
	fmt.Printf("Total Value: %v" , totalValue)
	fmt.Printf("Threes value %v", sumThree)
}
