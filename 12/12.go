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

type Rule struct{
	pattern string
	outcome string
}

func main() {
	f, err := os.Open("./test_input.txt")
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
	for line_no, line := range input {
		if line_no == 0 {
			parsed := regexp.MustCompile(`initial state: (.+$)`).FindStringSubmatch(line)
			state = strings.Split(parsed[1], "")
		} else {
			parsed := regexp.MustCompile(`([.#]+) => ([.#])`).FindStringSubmatch(line)
			rules = append(rules, Rule{parsed[1], parsed[2]})
		}
	}

	fmt.Println(strings.Join(state, ""))
	for i:=0; i<20; i++{
		step(&state, &rules)
		fmt.Println(strings.Join(state, ""))
	}
}

func step(state *[]string, rules *[]Rule){
	for _, rule := range *rules {
		for i:=0; i<len(*state); i++{ // i is current index into state
			matches := true
			for j, rule_letter := range rule.pattern {
				var state_letter string

				switch{
				case j == 0:
					if i - 2 < 0{
						state_letter = "."
					}else{
						state_letter = (*state)[i-2]
					}
				case j == 1:
					if i - 1 < 0{
						state_letter = "."
					}else{
						state_letter = (*state)[i-1]
					}
				case j == 2:
					state_letter = (*state)[i]
				case j == 3:
					if i + 1 > len(*state)-1{
						state_letter = "."
					}else{
						state_letter = (*state)[i+1]
					}
				case j == 4:
					if i + 2 > len(*state)-1{
						state_letter = "."
					}else{
						state_letter = (*state)[i+2]
					}
				}
				if string(rule_letter) != state_letter{
					matches = false
					break
				}
			}
			if matches {
				(*state)[i] = rule.outcome
			}
		}
	}
}
