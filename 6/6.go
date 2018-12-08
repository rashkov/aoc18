package main

import (
	"bufio"
	"fmt"
	"math"
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
	x, y   int
	symbol string
}

var coords []Coord

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

	coords = parse_coords(input)
	// fmt.Println(coords)

	width, height := find_extents(coords)
	width += 1
	height += 1
	// fmt.Println("width, height: ", width, height)

	var grid [][]string // This grid should be used in Y, X form
	grid = make([][]string, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]string, width)
		// for j := 0; j < width; j++ {
		// 	grid[i][j] = "-"
		// }
	}
	populate_grid(&grid, coords, width, height)
	find_largest_uncontained_area(grid, width, height)
	// print_grid(&grid)
}

func find_largest_uncontained_area(grid [][]string, width int, height int) {
	type Stat struct {
		size     int
		infinite bool
	}
	stats := make(map[string]Stat)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			sym := strings.ToLower(grid[y][x])
			st := stats[sym]
			size := st.size
			st.size = size + 1
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				st.infinite = true
			}
			stats[sym] = st
		}
	}

	var largest_sym string
	var largest_size int
	for candidate_sym, candidate_stat := range stats {
		if candidate_stat.infinite == false && candidate_stat.size > largest_size {
			largest_size = candidate_stat.size
			largest_sym = candidate_sym
		}
	}
	fmt.Println("Part 1: ", string(largest_sym[0]), largest_size)
}

func populate_grid(grid *[][]string, coords []Coord, width int, height int) {
	for _, coord := range coords {
		(*grid)[coord.y][coord.x] = coord.symbol
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point := (*grid)[y][x]
			if point == "" {
				// Set it to a lowercase version of its nearest coord, or "." if tied
				(*grid)[y][x] = "-"
				distances := make(map[int][]string)
				for _, coord := range coords {
					distance := math.Abs(float64(x-coord.x)) + math.Abs(float64(y-coord.y))
					// distances[coord.symbol] = int(distance)
					distances[int(distance)] = append(distances[int(distance)], coord.symbol)
				}

				var shortest_dist = width + height + 1
				for dist, _ := range distances {
					if dist < shortest_dist {
						shortest_dist = dist
					}
				}
				closest_symbols := distances[shortest_dist]
				if len(closest_symbols) > 1 {
					(*grid)[y][x] = "." // There is a tie
				} else {
					(*grid)[y][x] = string(int(closest_symbols[0][0]) + 32) // Make it lowercase
				}
			}
		}
	}
}

func print_grid(grid *[][]string) {
	fmt.Print("   ")
	for x := 0; x < len((*grid)[0]); x++ {
		fmt.Print(x, " ")
	}
	fmt.Print("\n")
	for y := 0; y < len(*grid); y++ {
		fmt.Println(y, (*grid)[y])
	}
}

func parse_coords(input []string) []Coord {
	for index, txt := range input {
		xy := strings.Split(txt, ", ")
		x64, err := strconv.ParseInt(xy[0], 10, 64)
		check(err)
		y64, err := strconv.ParseInt(xy[1], 10, 64)
		check(err)
		coords = append(coords, Coord{int(x64), int(y64), string(65 + index)})
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
	return max_x + 1, max_y + 1
}
