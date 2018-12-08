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

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

type Coord struct {
	x, y int
}

var coords []Coord

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

	coords = parse_coords(input)
	fmt.Println(coords)

	width, height := find_extents(coords)
	fmt.Println("width, height: ", width, height)

	// This grid should be used in Y, X form
	var grid [][]string
	grid = make([][]string, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		for j := 0; j < width; j++ {
			grid[i][j] = "-"
		}
	}
	populate_grid(&grid, coords)
	print_grid(&grid)
}

func populate_grid(grid *[][]string, coords []Coord) {
	for _, coord := range coords {
		(*grid)[coord.y][coord.x] = "*"
	}
}

func print_grid(grid *[][]string){
	fmt.Print("   ")
	for x := 0; x < len((*grid)[0]); x++{
		fmt.Print(x, " ")
	}
	fmt.Print("\n")
	for y := 0; y < len(*grid); y++{
		fmt.Println(y, (*grid)[y])
	}
}

func parse_coords(input []string) []Coord {
	for _, txt := range input {
		xy := strings.Split(txt, ", ")
		x64, err := strconv.ParseInt(xy[0], 10, 64)
		check(err)
		y64, err := strconv.ParseInt(xy[1], 10, 64)
		check(err)
		coords = append(coords, Coord{int(x64), int(y64)})
	}
	return coords
}

func find_extents(coords []Coord) (width, height int) {
	var (
		max_x = 0
		max_y = 0
	)
	for _, coord := range coords {
		if coord.x > max_x {
			max_x = coord.x
		}
		if coord.y > max_y {
			max_y = coord.y
		}
	}
	return max_x+1, max_y+1
}
