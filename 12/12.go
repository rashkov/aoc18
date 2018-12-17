package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

type Rule struct {
	pattern string
	outcome string
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()
	var input []string
	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}

	var state []string
	var rules []Rule
	var zero_place = 0
	for line_no, line := range input {
		if line_no == 0 {
			parsed := regexp.MustCompile(`initial state: (.+$)`).FindStringSubmatch(line)
			state = strings.Split(parsed[1], "")
		} else {
			parsed := regexp.MustCompile(`([.#]+) => ([.#])`).FindStringSubmatch(line)
			rules = append(rules, Rule{parsed[1], parsed[2]})
		}
	}

	//fmt.Println(strings.Join(state, ""))
	for i := 1; i <= 300; i++ {
		state = step(state, &rules, &zero_place)
		//fmt.Println(strings.Join(state, ""))
		fmt.Println(i, part_1_sum(state, zero_place))
	}
	fmt.Println("Part 1:", part_1_sum(state, zero_place))
	fmt.Println("Part 2:", "the score increases by 73 points per generation. Use y=mx+b to figure out the rest.")
	fmt.Println("Part 2:", (50000000000-300)*73+23676)

}

func part_1_sum(state[]string, offset int) int{
	var sum int
	for index, state_letter := range state {
		if string(state_letter) == "#"{
			sum += index - offset
		}
	}
	return sum
}

func step(state []string, rules *[]Rule, zero_place *int) []string {
	// Expand it if there's a flower within the first or last two pots
	if state[0] == "#" || state[1] == "#" {
		*zero_place += 2
		state = append([]string{".", "."}, state...)
	}
	if state[len(state)-1] == "#" || state[len(state)-2] == "#" {
		state = append(state, []string{".", "."}...)
	}
	new_state := make([]string, len(state))
	for i := 0; i < len(state); i++ { // i is current index into state
		found_rule := false
		for _, rule := range *rules {
			rule_matches := true
			for j, rule_letter := range rule.pattern {
				var state_letter string
				var offset = i - (2 - j)
				if offset < 0 || offset >= len(state) {
					state_letter = "."
				} else {
					state_letter = state[offset]
				}
				if string(rule_letter) != state_letter {
					rule_matches = false
				}
			}
			if rule_matches {
				//state[i] = rule.outcome
				new_state[i] = rule.outcome
				found_rule = true
				break
			} else {
				continue
			}
		}
		Use(found_rule)
		// if i == 7 {
		// 	fmt.Println("found a rule for 7?", found_rule)
		// 	fmt.Println((*state)[5:10])
		// }
		if !found_rule {
			//state[i] = "."
			new_state[i] = "."
		}
	}
	return new_state
}
