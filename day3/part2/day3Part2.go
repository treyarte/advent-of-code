package day3Part2

import (
	"fmt"
	"strings"
)

/**
Priority values need for each item a-z and A-Z
**/
func CreatePriorityValues() (map[rune]int){
	lowerCaseStr := "abcdefghijklmnopqrstuvwxyz"

	upperCaseStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	alphaChars := map[rune]int {}

	for index, char := range lowerCaseStr {
		alphaChars[char] = index + 1
	}

	for index, char := range upperCaseStr {
		alphaChars[char] = (index + 1) + len(lowerCaseStr)
	}

	return alphaChars
}

func GetGroupUniqueBadgeSum(listOfBackPacks []string) (int) {
	var uniqueItemsList [] string

	for _, backPackRow := range listOfBackPacks {
		uniqueItemsList = append(uniqueItemsList, RemoveDuplicates(backPackRow))
	}
	var total int

	alphaChars := CreatePriorityValues()
	
	uniqueObjs := map[rune]int{}

	fullListOfItems := [] map[rune]int{}

	for index, items := range uniqueItemsList {
		for _, item := range items {
			_, isInItems := uniqueObjs[item]

			if isInItems {
				uniqueObjs[item] += 1
			} else {					
					uniqueObjs[item] = 1
			}
		
		}
		fmt.Println(uniqueObjs)
		if index > 0 && (index + 1)%3 == 0 {
			fullListOfItems = append(fullListOfItems, uniqueObjs)
			uniqueObjs = map[rune]int{}
		}
	}

	for _, items := range fullListOfItems {
		for key, val := range items {
			if val >= 3 {				
				total += alphaChars[key]
			}
		}
	}

	return total
}

func RemoveDuplicates(items string) (string){
	uniqueItems := map[string]int {}

	var convertedStr string = ""

	for _, char := range strings.Split(items, "") {
		uniqueItems[char] = 0;
	}

	for key, _ := range uniqueItems {
		convertedStr += key
	}

	return convertedStr
}