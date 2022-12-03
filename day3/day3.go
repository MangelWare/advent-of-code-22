package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()

	fmt.Println("Part 2:")
	part2()
}

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

func get_charprio(c rune) int {
	c_code := int(c)
	if c_code < int('a') {
		// Uppercase
		return 27 + c_code - int('A')
	} else {
		// Lowercase
		return 1 + c_code - int('a')
	}
}

func get_double_item(line string) rune {
	for i := 0; i < len(line)/2; i++ {
		if strings.ContainsRune(line[len(line)/2:], []rune(line)[i]) {
			return []rune(line)[i]
		}
	}
	panic("Nothing found")
}

func part1() {
	// Test for get_charprio
	/*
		for i := 'A'; i <= 'z'; i++ {
			fmt.Printf("%c: %d\n", i, get_charprio(i))
		}
	*/

	// Test for string splicing
	/*
		teststr := "AAABBB"
		fmt.Println(teststr[len(teststr)/2:])
	*/

	input := read_input()

	res := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		doubled_item := get_double_item(line)
		res += get_charprio(doubled_item)
	}

	fmt.Printf("Result: %d\n", res)
}

func find_badge(group [3]string) rune {
	for i := 0; i < len(group[0]); i++ {
		curr_char := string(group[0][i])
		if strings.Contains(group[1], curr_char) && strings.Contains(group[2], curr_char) {
			return rune(group[0][i])
		}
	}
	panic("No badge found")
}

func part2() {
	input := read_input()

	sum := 0

	group_lines := [3]string{"", "", ""}
	for line_idx, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		group_idx := line_idx % 3
		group_lines[group_idx] = line
		if group_idx == 2 {
			// Group is complete, calculate result
			badge := find_badge(group_lines)
			sum += get_charprio(badge)
		}
	}

	fmt.Printf("Result: %d\n", sum)
}
