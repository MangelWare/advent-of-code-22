package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Enemy
A Rock
B Paper
C Scissors

Self
X Rock
Y Paper
Z Scissors

Score for round:
Shape score + outcome score
Shape: 1 Rock, 2 Paper, 3 Scissors
Outcome: 0 lose, 3 draw, 6 win
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_input() string {
	dat, err := os.ReadFile("./input")
	check(err)
	return string(dat)
}

func main() {

	//puzzle_version := 1
	puzzle_version := 2

	input := read_input()

	total_score := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		enemy_choice := line[0]
		own_choice := line[1]
		if puzzle_version == 2 {
			own_choice = get_choice_ver2(enemy_choice, line[1])
		}

		ss := shape_score(own_choice)
		os := outcome_score(enemy_choice, own_choice)
		total_score += ss
		total_score += os

		// For debugging
		//fmt.Printf("Playing %s against %s, Score is %d (Shape) + %d (Outcome)\n", conv_shape(line[1]), conv_shape(line[0]), ss, os)
	}

	fmt.Printf("Total score: %d", total_score)
}

func shape_score(choice string) int {
	switch choice {
	case "X", "A":
		return 1 // Rock
	case "Y", "B":
		return 2 // Paper
	case "Z", "C":
		return 3 // Scissors
	}
	panic(choice)
}

func conv_shape(in string) string {
	switch in {
	case "X", "A":
		return "Rock"
	case "Y", "B":
		return "Paper"
	case "Z", "C":
		return "Scissors"
	}
	panic(in)
}

func outcome_score(enemy_choice string, own_choice string) int {
	// We reuse shape scores for decision making
	ec_idx := shape_score(enemy_choice) % 3
	oc_idx := shape_score(own_choice) % 3
	switch (ec_idx - oc_idx + 3) % 3 {
	case 0:
		return 3 // draw
	case 1:
		return 0 // lose
	case 2:
		return 6 // win
	}
	panic(enemy_choice + " " + own_choice)
}

func get_choice_ver2(enemy_choice string, result string) string {
	// Convert ver2 to ver1 input
	idx := shape_score(enemy_choice)

	switch result {
	case "X": // lose
		idx -= 1
	case "Z": // win
		idx += 1
	}

	// Convert back to X,Y,Z
	switch (idx + 3) % 3 {
	case 0:
		return "Z"
	case 1:
		return "X"
	case 2:
		return "Y"
	}

	panic("AAAAAAH!")
}
