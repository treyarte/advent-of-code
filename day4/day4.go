package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 - Loop through the file
- add both pairs as one object in an array
- the obj props are pair1 and pair2
- they are arrays of max 2
- ex: [
		{
			pair1: [18, 20],
			pair2: [19, 21]
		},
		{
			pair1: [9, 86]
			pair2: [9, 87]
		}
	]
- loop through our new arr of objs
- for each object take the first index of each property
- compare the numbers
- if one is less than the other then we will focus on the opposite side
- ex: 18 is less than 19 so we focus of pair2
- if greater than, then focus on the other pair
- if equal will start with the first pair
		- check if the 2nd pair 2nd index fit into is less than or equal to the first pair 2nd index
		- if so we have a pair that fits
		- if the 2nd pair is greater than we still have one that fits because their first value are equal to each other
		- in a sense if their first value is equal to each other we are guaranteed one that fits
- then take the 2nd index and check if it is less than pair2 2nd index val
- if it is less then we have a pair that fits into one
- if it is greater we don't
-if the are equal we have a fit
*/

var pl = fmt.Println

type pairOfElves struct {
	pair1 [2]int	
	pair2 [2]int	
}

func main() {
	var arrOfElfPairs []pairOfElves

	var numOfFits = 0;

	file, err := os.Open("./day_4_input.txt");
	
	if err != nil {
		panic(err);
	}
	defer file.Close()
	

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		text := scanner.Text()
		spl := strings.Split(text, ",")
		spl1 := strings.Split(spl[0], "-")
		spl2 := strings.Split(spl[1], "-")
		
		pairObj := pairOfElves {
			pair1: ConvertArrStrToIntArr(spl1),
			pair2: ConvertArrStrToIntArr(spl2),
		}
		arrOfElfPairs = append(arrOfElfPairs, pairObj)
	}

	numOfFits = GetNumOfPairFits(arrOfElfPairs);

	numberOfOverlaps := GetNumOfOverlaps(arrOfElfPairs)

	pl(numberOfOverlaps)
	
	pl(numOfFits)
}

/**
	- Take the pair 1
	- if pair 1 index is equal to pair 2 index 1
		- we have an overlap
	- if it is greater than pair 2
		- focus on pair 1
		- is pair1 index 1 greater than pair2 index 2
			- if equal we have an overlap
			- if less than we do not have a match
			- if greater than we have overlap
	- focus on pair 2
	 - take pair 1 index 2
	 - check if pair 1 index 2 is less than pair 2 index 1
	 	- if equal to then we have an overlap
		- if less than we do not have an overlap
		- if Greater than then check index 2 of pair 2
			- if greater than index two we dont have a match
			- if less than then we have a match
	-
	
**/
func GetNumOfOverlaps(arrOfElfPairs []pairOfElves)(int) {
	var numOfOverlaps int = 0

	for _, val := range arrOfElfPairs {
		pair1 := val.pair1
		pair2 := val.pair2

		if pair1[0] == pair2[0] {
			numOfOverlaps++
			continue
		}

		if pair1[0] > pair2[0] {
			if pair1[0] <= pair2[1] {
				numOfOverlaps++
				continue
			}
		}

		if pair1[0] < pair2[0] {
			if pair1[1] == pair2[0] {
				numOfOverlaps++
				continue
			}

			if pair1[1] > pair2[0] {
				numOfOverlaps++
				continue
			}
		}
	}
	return numOfOverlaps
}

func GetNumOfPairFits(arrOfElfPairs []pairOfElves) (int) {
	var numOfFits int = 0

	for _, val := range arrOfElfPairs {
		pair1 := val.pair1
		pair2 := val.pair2

		if pair1[0] == pair2[0] {
			numOfFits += 1
			continue
		}

		if pair1[0] < pair2[0] {
			if pair1[1] >= pair2[1] {
				numOfFits+= 1
				continue
			}
		}

		if pair1[0] > pair2[0] {
			if pair1[1] <= pair2[1] {
				numOfFits+= 1
				continue
			}
		}
	}

	return numOfFits
}

func ConvertArrStrToIntArr(strArr []string) ([2] int) {
	converted1, _ := strconv.ParseInt(strArr[0], 10, 64)
	converted2, _ := strconv.ParseInt(strArr[1], 10, 64)
	
	arrInt := [2]int{int(converted1), int(converted2)}	
	return arrInt
}