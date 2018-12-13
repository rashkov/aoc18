package main

import (
	"fmt"
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

func main() {
	serial_num := 6303 // from puzzle input

	var grid [][]int = make([][]int, 301)
	for i:=0; i<301; i++{
		grid[i] = make([]int, 301)
	}

	var cached_power_level [][][]int = make([][][]int, 301)
	for i:=0; i<301; i++{
		cached_power_level[i] = make([][]int, 301)
		for j:=0; j<301; j++{
			cached_power_level[i][j] = make([]int, 301)
			for k:=0; k<301; k++{
				cached_power_level[i][j][k] = -999
			}
		}
	}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			//x_str := strconv.FormatInt(int64(x), 10)
			pwr := calc_power_level(x, y, serial_num)
			grid[y][x] = pwr
		}
		//y_str := strconv.FormatInt(int64(y), 10)
		//fmt.Println(y_str, grid[y])
	}
	var (
		max_power int
		max_x     int
		max_y     int
		max_side  int
	)
	max_power, max_x, max_y = calc_max_power_level(&grid, 3, &cached_power_level)
	fmt.Println("Part 1", "power:", max_power, "x:", max_x, "y:", max_y)

	max_power = -1
	max_x = -1
	max_y = -1
	max_side = -1
	for side := 1; side <= 300; side++ {
		power, x, y := calc_max_power_level(&grid, side, &cached_power_level)
		if power > max_power{
			max_power = power
			max_x = x
			max_y = y
			max_side = side
		}
	}
	fmt.Println("Part 2", "power:", max_power, "x:", max_x, "y:", max_y, "side:", max_side)

}

func calc_max_power_level(grid *[][]int, side int, cached_power_level *[][][]int) (int, int, int){
	var (
		max_power int
		max_x     int
		max_y     int
	)
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			if (*cached_power_level)[y][x][side-1] == -999{
				// x,y is upper-left corner of our box
				if x+side-1 <= 300 && y+side-1 <= 300 {
					var total_power int
					for k := y; k < y+side; k++ {
						for j := x; j < x+side; j++ {
							total_power += (*grid)[k][j]
						}
					}
					if total_power > max_power {
						max_power = total_power
						max_x = x
						max_y = y
					}
					(*cached_power_level)[y][x][side] = total_power
				}
			}else{
				if x+side-1 <= 300 && y+side-1 <= 300 {
					total_power := (*cached_power_level)[y][x][side-1]
					// add row with y = y + side - 1
					K := y + side - 1
					for j := x; j < x+side; j++{
						total_power += (*grid)[K][j]
					}
					// add the column with x = x + side - 1
					L := x + side - 1
					for m := y; m < y + side; m++{
						total_power += (*grid)[m][L]
					}
					if total_power > max_power {
						max_power = total_power
						max_x = x
						max_y = y
					}
					(*cached_power_level)[y][x][side] = total_power
				}
			}
		}
	}
	return max_power, max_x, max_y
}

func calc_power_level(x_coord int, y_coord int, grid_serial_number int) int {
	var (
		power_level int
		rack_id     int
	)

	rack_id = x_coord + 10
	power_level = hundreds_digit_only(float64((rack_id*y_coord+grid_serial_number)*rack_id)) - 5
	return power_level
}

func hundreds_digit_only(x float64) int {
	// i is us how many digits the input number is
	str := strconv.FormatFloat(x, 'f', 0, 64)
	hundreds_place := len(str) - 3
	if hundreds_place < 0 {
		return 0
	} else {
		parsed, err := strconv.ParseInt(string(str[hundreds_place]), 10, 64)
		check(err)
		return int(parsed)
	}
}

func tests() {
	// This tests hundreds_digit_only()
	for _, n := range []int{5, 12, 99, 153, 1396, 19223, 199333} {
		fmt.Println("hundreds digit", n, "is:", hundreds_digit_only(float64(n)))
	}

	// This tests against test input
	type Input struct {
		x      int
		y      int
		serial int
		level  int
	}
	for _, input := range []Input{Input{217, 196, 39, 0}, Input{122, 79, 57, -5}, Input{101, 153, 71, 4}} {
		pwr := calc_power_level(input.x, input.y, input.serial)
		fmt.Println(input, pwr)
	}
}
