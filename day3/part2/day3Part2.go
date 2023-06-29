package day3Part2

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

func GetGroupUniqueBadgeSum(listOfBackPacks []string) {
	
}