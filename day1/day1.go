package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sumSlice(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Day 1 has arrived!")
	var inputs []string
	var results []int

	// Load input value
	// Separate lines into array of strings
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file")
		return
	}

	fmt.Printf("Inputs: %d\n", len(inputs))

	// Loop through each string
	// -- Find the first digit from the front
	// -- Find the last digit from the back
	// -- Concatenate them together as int
	regexFront, err := regexp.Compile(`^\D*(\d)`)
	if err != nil {
		fmt.Println("Error compiling regex for forward digit search")
		return
	}
	regexBack, err := regexp.Compile(`(\d)\D*$`)
	if err != nil {
		fmt.Println("Error compiling regex for backward digit search")
		return
	}

	for index, input := range inputs {
		fmt.Printf("Input %d: %s\n", index+1, input)
		front := regexFront.FindStringSubmatch(input)
		back := regexBack.FindStringSubmatch(input)
		resultInt, err := strconv.Atoi(front[1] + back[1])
		if err != nil {
			fmt.Println("Error converting resulting string to int")
			return
		}
		fmt.Printf("Result: %d\n", resultInt)
		results = append(results, resultInt)
	}

	// Sum all values in the array
	fmt.Printf("\nSum: %d\n", sumSlice(results))
}
