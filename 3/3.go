package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		coords := txt[2]
		extents := txt[3]

		coordRE := regexp.MustCompile(`(\d+),(\d+)`)
		fmt.Println(coordRE.FindStringSubmatch(coords))

		extentsRE := regexp.MustCompile(`(\d+)x(\d+)`)
		fmt.Println(extentsRE.FindStringSubmatch(extents))
	}
	check(scanner.Err())
}
