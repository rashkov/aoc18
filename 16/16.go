package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type RunSet struct {
	before      [4]int
	instruction [4]int
	after       [4]int
}

type RunSets []RunSet

func (runset *RunSet) complete() bool {
	return runset.before[0] != -1 && runset.instruction[0] != -1 && runset.after[0] != -1
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var runsets RunSets
	run_set := RunSet{[4]int{-1, 0, 0, 0}, [4]int{-1, 0, 0, 0}, [4]int{-1, 0, 0, 0}}

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			continue
		}
		if run_set.complete() {
			runsets = append(runsets, run_set)
			run_set = RunSet{[4]int{-1, 0, 0, 0}, [4]int{-1, 0, 0, 0}, [4]int{-1, 0, 0, 0}}
		}
		if match, _ := regexp.MatchString("Before", txt); match {
			parsed := regexp.MustCompile(`(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)`).FindStringSubmatch(txt)
			parsed = parsed[1:]
			// fmt.Println(txt)
			// fmt.Println(parsed[1:])
			// before := parsed[1:]
			//run_set.before = parsed[1:5]
			for i := 0; i < 4; i++ {
				parsed_int64, err := strconv.ParseInt(parsed[i], 10, 64)
				check(err)
				run_set.before[i] = int(parsed_int64)
			}
		}
		if match, _ := regexp.MatchString("After", txt); match {
			parsed := regexp.MustCompile(`(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)`).FindStringSubmatch(txt)
			// fmt.Println(txt)
			// fmt.Println(parsed[1:])
			parsed = parsed[1:]
			// after := parsed[1:]
			// run_set.after = parsed[1:]
			for i := 0; i < 4; i++ {
				parsed_int64, err := strconv.ParseInt(parsed[i], 10, 64)
				check(err)
				run_set.after[i] = int(parsed_int64)
			}
		}
		if match, _ := regexp.MatchString(`^(\d+)[, ]`, txt); match {
			parsed := regexp.MustCompile(`(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)[, ]{1,}(\d+)`).FindStringSubmatch(txt)
			parsed = parsed[1:]
			// fmt.Println(txt)
			// fmt.Println(parsed[1:])
			// instruction := parsed[1:]
			// run_set.instruction = instruction
			for i := 0; i < 4; i++ {
				parsed_int64, err := strconv.ParseInt(parsed[i], 10, 64)
				check(err)
				run_set.instruction[i] = int(parsed_int64)
			}
		}
	}
	fmt.Println(runsets)

}
