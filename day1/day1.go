package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("./input")
	check(err)

	input_str := string(dat)

	scanner := bufio.NewScanner(strings.NewReader(input_str))
	curr_calories := 0
	max_calories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			c, err := strconv.Atoi(line)
			check(err)
			curr_calories += c
		} else {
			if curr_calories > max_calories {
				max_calories = curr_calories
			}
			curr_calories = 0
		}
	}

	fmt.Printf("Max: %d", max_calories)

}
