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

	var matx [1000][1000] int
	for i := 0; i < 1000; i++ {
		for j := 0; i < 1000; i++ {
			matx[i][j] = 0
		}
	}

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		coords := txt[2]
		extents := txt[3]

		coordRE := regexp.MustCompile(`(\d+),(\d+)`)
		parsed_coords := coordRE.FindStringSubmatch(coords)
		x_coord, err := strconv.ParseInt(parsed_coords[1], 10, 64)
		check(err)
		y_coord, err := strconv.ParseInt(parsed_coords[2], 10, 64)
		check(err)

		extentsRE := regexp.MustCompile(`(\d+)x(\d+)`)
		parsed_extents := extentsRE.FindStringSubmatch(extents)
		x_extent, err := strconv.ParseInt(parsed_extents[1], 10, 64)
		check(err)
		y_extent, err := strconv.ParseInt(parsed_extents[2], 10, 64)
		check(err)

		// fmt.Println("x: ", x_coord, "y: ", y_coord)
		// fmt.Println("width: ", x_extent, "height: ", y_extent)
		for i := int64(0); i < x_extent; i++ {
			for j := int64(0); j < y_extent; j++ {
				matx[x_coord + i][y_coord + j] += 1
			}
		}
	}

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if matx[i][j] > 1{
				count++
			}
		}
	}
	fmt.Println("Part 1: ", count)
}
