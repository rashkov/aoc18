package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

type Point struct {
	x  int64
	y  int64
	dx int64
	dy int64
}

var points []*Point

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
	for _, line := range input {
		//fmt.Println(line)
		//parsed := regexp.MustCompile(`position=<.?([-\s]\d+)`).FindStringSubmatch(line)
		parsed := regexp.MustCompile(`position=<(.*?),(.*?)>.*?<(.*?),(.*?)>`).FindStringSubmatch(line)
		x, err := strconv.ParseInt(strings.TrimSpace(parsed[1]), 10, 64)
		check(err)
		y, err := strconv.ParseInt(strings.TrimSpace(parsed[2]), 10, 64)
		check(err)
		dx, err := strconv.ParseInt(strings.TrimSpace(parsed[3]), 10, 64)
		check(err)
		dy, err := strconv.ParseInt(strings.TrimSpace(parsed[4]), 10, 64)
		check(err)
		//fmt.Println(x, y, dx, dy)
		points = append(points, &Point{x, y, dx, dy})
	}

	var n int
	var largest_x_count, largest_y_count int64
	for n = 0; n < 10036; n++ {
		step(&points)
		var found_new_x = false
		var found_new_y = false
		x, x_count := check_if_points_align(&points, "x")
		if x_count > largest_x_count {
			found_new_x = true
			largest_x_count = x_count
			fmt.Println(x_count, "points had x value", x, "at n=", n)
		}

		y, y_count := check_if_points_align(&points, "y")
		if y_count > largest_y_count {
			found_new_y = true
			largest_y_count = y_count
			fmt.Println(y_count, "points had y value", y, "at n=", n)
		}
		if found_new_x && found_new_y && n != 0 {
			fmt.Println("Found intersection point at n=", n)
			break
		}
	}

	smallest_x, smallest_y, largest_x, largest_y := get_range(&points)
	fmt.Println("Range of:")
	fmt.Println(smallest_x, smallest_y, largest_x, largest_y)

	points_index := index_points(&points)
	// fmt.Println(points_index)

	fmt.Println("Part 1:")
	for y := smallest_y; y <= largest_y; y++ {
		var out []string
		for x := smallest_x; x <= largest_x; x++ {
			if points_index[y][x] != nil {
				out = append(out, "#")
			}else{
				out = append(out, ".")
			}
		}
		fmt.Println(strings.Join(out, ""))
	}

	fmt.Println("Part 2:", n+1)
}

// See if some fraction of the total points align in the x-axis
func check_if_points_align(points *[]*Point, axis string) (int64, int64) {
	stats := make(map[int64]int64)
	for _, point_ptr := range *points {
		var coord int64
		if axis == "x" {
			coord = (*point_ptr).x
		} else {
			coord = (*point_ptr).y
		}
		stats[coord] += 1
	}

	var max, count int64
	for coord, stat := range stats {
		if coord > max {
			max = coord
			count = stat
		}
	}
	return max, count
}

func step(points *[]*Point) {
	for _, point_ptr := range *points {
		(*point_ptr).x = (*point_ptr).x + (*point_ptr).dx
		(*point_ptr).y = (*point_ptr).y + (*point_ptr).dy
	}
}

func get_range(points *[]*Point) (int64, int64, int64, int64) {
	var smallest_x, smallest_y, largest_x, largest_y int64
	first_point := (*points)[0]
	smallest_x = (*first_point).x
	smallest_y = (*first_point).y
	largest_x = (*first_point).x
	largest_y = (*first_point).y

	for _, point_ptr := range *points {
		if (*point_ptr).x < smallest_x {
			smallest_x = (*point_ptr).x
		}
		if (*point_ptr).y < smallest_y {
			smallest_y = (*point_ptr).y
		}
		if (*point_ptr).x > largest_x {
			largest_x = (*point_ptr).x
		}
		if (*point_ptr).y > largest_y {
			largest_y = (*point_ptr).y
		}
	}
	return smallest_x, smallest_y, largest_x, largest_y
}

func index_points(points *[]*Point) map[int64](map[int64](*Point)) {
	index := make(map[int64](map[int64](*Point)))
	for _, point_ptr := range *points {
		if index[(*point_ptr).y] == nil {
			index[(*point_ptr).y] = make(map[int64](*Point))
		}
		index[(*point_ptr).y][(*point_ptr).x] = point_ptr
	}
	return index
}
