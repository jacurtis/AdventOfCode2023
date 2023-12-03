package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sumSlice(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return sum
}

func convStrNumber(str string) string {
	num, err := strconv.Atoi(str)
	if err != nil {
		switch str {
		case "one":
			num = 1
		case "two":
			num = 2
		case "three":
			num = 3
		case "four":
			num = 4
		case "five":
			num = 5
		case "six":
			num = 6
		case "seven":
			num = 7
		case "eight":
			num = 8
		case "nine":
			num = 9
		}
	}
	representation := strconv.Itoa(num)
	return representation

}

func Part1(debug ...bool) {
	if len(debug) == 0 {
		debug = append(debug, false)
	}

	fmt.Println("Day 1: Challenge 1")
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

	if debug[0] {
		fmt.Printf("Inputs: %d\n", len(inputs))
	}

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
		if debug[0] {
			fmt.Printf("Input %d: %s\n", index+1, input)
		}
		front := regexFront.FindStringSubmatch(input)
		back := regexBack.FindStringSubmatch(input)
		resultInt, err := strconv.Atoi(front[1] + back[1])
		if err != nil {
			fmt.Println("Error converting resulting string to int")
			return
		}
		if debug[0] {
			fmt.Printf("Result: %d\n", resultInt)
		}
		results = append(results, resultInt)
	}

	// Sum all values in the array
	fmt.Printf("Sum: %d\n", sumSlice(results))
}

func Part2(debug ...bool) {
	fmt.Println("Day 1: Challenge 2")
	if len(debug) == 0 {
		debug = append(debug, false)
	}
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

	if debug[0] {
		fmt.Printf("Inputs: %d\n", len(inputs))
	}

	// Loop through each string
	// -- Find the first digit from the front
	// -- Find the last digit from the back
	// -- Concatenate them together as int
	r, err := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	if err != nil {
		fmt.Println("Error compiling regex for forward digit search")
		return
	}

	for index, input := range inputs {
		if debug[0] {
			fmt.Printf("Input %d: %s\n", index+1, input)
		}
		//matches := r.FindAllString(input, -1)
		var matches []string
		for i := range input {
			match := r.FindString(input[i:])
			if match != "" {
				matches = append(matches, match)
			}
		}
		resultInt, err := strconv.Atoi(convStrNumber(matches[0]) + convStrNumber(matches[len(matches)-1]))
		if err != nil {
			fmt.Println("Error converting resulting string to int")
			return
		}
		if debug[0] {
			fmt.Printf("Result: %d\n", resultInt)
		}
		results = append(results, resultInt)
	}

	// Sum all values in the array
	fmt.Printf("Sum: %d\n", sumSlice(results))
}

func main() {
	Part1()
	fmt.Println(strings.Repeat("=", 20))
	Part2()
}
