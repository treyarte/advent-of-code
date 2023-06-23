package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

/*
	Create a var to hold max value
	create a var to hold our currentCalTotal
	read through the file
	for each line add up the calories
	if we hit an empty line check if the added calories are greater than the current max
	if they are greater than the max update the max
	if not update current count to 0 and repeat until end of the file
*/

var pl = fmt.Println

func main() {
	var maxCalories float64
	var currentCalCount float64
	maxCalories = 0
	currentCalCount = 0

	
	//opening the file
	file, err := os.Open("adventofcode_2022_day_1_input.txt")
	if err != nil {
		log.Fatal("Error opening the file", err)
		return
	} 
	defer file.Close() //ensure the file is closed at the end

	//creates a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maxCalories = math.Max(maxCalories, currentCalCount)
			currentCalCount = 0;
		} else {
			intVal, err := strconv.Atoi(line); 

			if err != nil {
				pl("Error converting string to int", err)
				continue
			}
		
			currentCalCount += float64(intVal)
			pl(maxCalories)
		}




		
		// pl(currentCalCount)
		// pl(line)
	}

	if err := scanner.Err(); err != nil {
		pl("Error reading file:", err)
	}

}