package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var input []string
	data_struct := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		// txt := strings.Split(scanner.Text(), " ")
		input := append(input, scanner.Text())

		// coordRE := regexp.MustCompile(`(\d+),(\d+)`)
		// parsed_coords := coordRE.FindStringSubmatch(coords)
		// x_coord, err := strconv.ParseInt(parsed_coords[1], 10, 64)
		// check(err)

	}
	fmt.Println(input)
}
