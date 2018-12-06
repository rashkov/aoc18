package main

import (
	"bufio"
	"fmt"
	"os"
	_ "regexp"
	"sort"
	_ "strconv"
	"unsafe"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Guard struct {
	guard_id     int
	total_asleep int
}

func main() {
	var a int
	fmt.Println(unsafe.Sizeof(a))
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var input []string
	// guards := make(map[int]Guard)
	// var current_guard, awoke_minute, asleep_minute int

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	sort.Strings(input)
	for _, txt := range input {
		// if match, _ := regexp.MatchString("Guard", txt); match {
		// 	parsed := regexp.MustCompile(`#(\d+)`).FindStringSubmatch(txt)
		// 	guard_id, err := strconv.ParseInt(parsed[1], 10, 8)
		// 	check(err)
		// 	fmt.Println("Guard!", guard_id)
		// 	current_guard = guard_id
		// } else if match, _ := regexp.MatchString("asleep", txt); match {
		// 	parsed := regexp.MustCompile(`\d+:(\d+)`).FindStringSubmatch(txt)
		// 	asleep_minute, err := strconv.ParseInt(parsed[1], 10, 64)
		// 	check(err)
		// 	fmt.Println("asleep", asleep_minute)
		// } else if match, _ := regexp.MatchString("wakes", txt); match {
		// 	parsed := regexp.MustCompile(`\d+:(\d+)`).FindStringSubmatch(txt)
		// 	awoke_minute, err := strconv.ParseInt(parsed[1], 10, 64)
		// 	check(err)
		// 	fmt.Println("awake", minute)
		// 	guards[current_guard].total_asleep += awoke_minute - asleep_minute
		// }
	}
}

//var txt []string
// data_struct := make(map[string]bool)
// txt = strings.Split(scanner.Text(), "\n")
// actionRE := regexp.MustCompile(`#(\d+)`)
// actionRE.FindStringSubmat(in:(\d+)ut)// .ParseInt(parsed_coords[1], 10, 64)
