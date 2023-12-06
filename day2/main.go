package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var game_rounds = make(map[int][][]string)
var max = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}
var total = 0

func parseGame(id int, str string) {
	//fmt.Println(id, str)

	r := regexp.MustCompile(`^Game\s(\d+):\s(.*)$`)
	matches := r.FindStringSubmatch(str)
	//fmt.Println(matches[1], matches[2])
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println("Error converting string to int")
		os.Exit(1)
	}

	var game_arr [][]string
	split := strings.Split(matches[2], ";")
	for _, val := range split {
		indv_games := strings.Split(val, ",")
		game_arr = append(game_arr, indv_games)
	}
	game_rounds[id] = game_arr
}

func Part1(debug ...bool) {
	// Set debug to false if not provided so it is safe to reference later
	if len(debug) == 0 {
		debug = append(debug, false)
	}

	// Open File
	input, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parseGame(0, scanner.Text())
	}
	//
	for round_number, game := range game_rounds {
		valid := true
		for _, v := range game {
			for _, item := range v {
				parsed := strings.Split(strings.TrimSpace(item), " ")
				cube := parsed[1]
				num, _ := strconv.Atoi(parsed[0])

				if num <= max[cube] {
					fmt.Println("Valid", cube, num)
				} else {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			total += round_number
		}
	}
	fmt.Println("Total: ", total)
}

func main() {
	fmt.Println("Running Day 2")

	Part1()

	//fmt.Println(game_rounds)
}
