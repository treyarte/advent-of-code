package main

import (
	day3Part2 "artetech/part2"
	"bufio"
	"fmt"
	"os"
)

var pl = fmt.Println

/**
Given a text file of bag items per bag
find the total sum of items that are in both compartments of each bag
each line in the text file is considered a bag. The first half of characters are items that
belong to the first compartment. While the second half of characters belong to the other.
- characters are case sensitive, i.e: z and Z are different items
- a - z = 1 - 25 and A - Z = 27 -52

steps:
- read file
- while reading the file
- store each line as an item in an array
- loop through the array
- Add the string to a unique object until a counter is half of the string length
- check the unique object against the rest of the string
- if the next half is in the unique object add += 1
- After loop is over get an array of unique items in object whose value is greater than 0
- look through that array and use the a-z and A-Z objects to get the total

**/

func checkError(err error) {
	if(err != nil) {
		panic(err)
	}
}

func main() {

	var listOfBackPacks []string

	file, err := os.Open("adventofcode.com_2022_day_3_input.txt");
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	//not needed because this is default
	// scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text :=  scanner.Text()
		listOfBackPacks	= append(listOfBackPacks,text)
	}

	alphaChars := day3Part2.CreatePriorityValues()

	var total int = 0

	for _, backpack := range listOfBackPacks {
		counter := 1
		counterMax := len(backpack)/2
		uniqueItems := map[rune]int {}
		for _, item := range backpack {
			if counter <= counterMax {
				uniqueItems[item] = 0
				counter++
				continue
			}

			if  val, existInUnique := uniqueItems[item]; existInUnique == true && val <= 0{				
				uniqueItems[item] += 1
				total+= alphaChars[item]
			} 
		}
	}
	pl(total)

}