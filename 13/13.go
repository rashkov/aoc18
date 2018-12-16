package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

type Cart struct {
	sym       string
	x         int
	y         int
	prev_x    int
	prev_y    int
	num_turns int
	grid_ptr  *[][]string
	carts_ptr *Carts
}

// y, x, array of pointers to cart. if len(array) > 1, then carts crashed
type Carts [][][]*Cart

var carts Carts

var part_1_complete = false

func print(grid_ptr *[][]string, carts_ptr *Carts) {
	grid := *grid_ptr
	carts := *carts_ptr
	for y, row := range grid {
		fmt.Print(y)
		for x, col := range row {
			if len(carts[y][x]) == 1 {
				if carts[y][x][0] != nil {
					fmt.Print((*carts[y][x][0]).sym)
				}
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func (cart_ptr *Cart) move(next_x, next_y int) {
	(*cart_ptr).prev_x = (*cart_ptr).x
	(*cart_ptr).prev_y = (*cart_ptr).y
	(*cart_ptr).x = next_x
	(*cart_ptr).y = next_y
	(*cart_ptr).assert_on_track()

	carts_ptr := (*cart_ptr).carts_ptr

	carts_at_xy := (*carts_ptr)[cart_ptr.prev_y][cart_ptr.prev_x]
	for i, cart_at_xy_ptr := range carts_at_xy {
		if cart_at_xy_ptr == cart_ptr {
			// Remove it from its old location
			array_delete(&carts_at_xy, i)
			// Don't forget to assign it back!
			(*carts_ptr)[cart_ptr.prev_y][cart_ptr.prev_x] = carts_at_xy

			// insert it into its new location
			(*carts_ptr)[cart_ptr.y][cart_ptr.x] = append((*carts_ptr)[cart_ptr.y][cart_ptr.x], cart_ptr)

			// Check crash condition
			if len((*carts_ptr)[cart_ptr.y][cart_ptr.x]) > 1 {
				if(!part_1_complete){
					part_1_complete = true
					crash_str := fmt.Sprintf("The carts have CRASHED! at x=%d, y=%d", cart_ptr.x, cart_ptr.y)
					fmt.Println("Part 1:", crash_str)
				}

				// Remove the crashed carts from the carts collection
				(*carts_ptr)[cart_ptr.y][cart_ptr.x] = []*Cart{}

				// If there's only one cart left, print its location and exit
				var carts_remaining []*Cart
				for _, row := range *carts_ptr {
					for _, col := range row {
						if len(col) > 0 {
							carts_remaining = append(carts_remaining, col...)
						}
					}
				}
				fmt.Println("Carts remaining:", len(carts_remaining))
				if len(carts_remaining) == 1 {
					last_cart := (*carts_remaining[0])
					fmt.Printf("Part 2: Last cart remaining at: x=%d, y=%d", last_cart.x, last_cart.y)
					os.Exit(0)
				}
			}
		}
	}
}

func (cart_ptr *Cart) assert_on_track() {
	x, y := cart_ptr.x, cart_ptr.y
	grid_ptr := cart_ptr.grid_ptr
	grid_sym := (*grid_ptr)[y][x]
	if grid_sym == " " {
		fmt.Println("Jumped off the track")
		fmt.Println(cart_ptr)
		panic("Jumped off the track")
	}
}

func (cart_ptr *Cart) turn(direction string) {
	var next_cart_sym string
	if direction == "left" {
		switch cart_ptr.sym {
		case ">":
			next_cart_sym = "^"
		case "<":
			next_cart_sym = "v"
		case "^":
			next_cart_sym = "<"
		case "v":
			next_cart_sym = ">"
		}
	} else if direction == "right" {
		switch cart_ptr.sym {
		case ">":
			next_cart_sym = "v"
		case "<":
			next_cart_sym = "^"
		case "^":
			next_cart_sym = ">"
		case "v":
			next_cart_sym = "<"
		}
	} else {
		panic("Invalid direction passed to turn()")
	}
	(*cart_ptr).sym = next_cart_sym
}

func (cart_ptr *Cart) step() {
	next_x, next_y, next_track_sym := cart_ptr.get_next_coord()

	switch next_track_sym {
	case `+`:
		// turn, use cart's turning algorithm
		turns_mod := int(math.Mod(float64(cart_ptr.num_turns), float64(3)))
		if turns_mod == 0 {
			(*cart_ptr).turn("left")
		} else if turns_mod == 1 {
			// go straight: don't change the symbol
		} else if turns_mod == 2 {
			// go right
			(*cart_ptr).turn("right")
		}
		(*cart_ptr).num_turns += 1
	case `/`:
		// turn cart based on its approach direction
		switch cart_ptr.sym {
		case ">":
			(*cart_ptr).turn("left")
		case "<":
			(*cart_ptr).turn("left")
		case "^":
			(*cart_ptr).turn("right")
		case "v":
			(*cart_ptr).turn("right")
		}
	case `\`:
		// turn cart based on its approach direction
		switch cart_ptr.sym {
		case ">":
			(*cart_ptr).turn("right")
		case "<":
			(*cart_ptr).turn("right")
		case "^":
			(*cart_ptr).turn("left")
		case "v":
			(*cart_ptr).turn("left")
		}
	}
	// Move the cart to the new coordinate
	(*cart_ptr).move(next_x, next_y)
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()
	var grid [][]string
	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	carts = make([][][]*Cart, len(grid))
	for y, row := range grid {
		if carts[y] == nil {
			carts[y] = make([][]*Cart, len(row))
		}
		for x, sym_rune := range row {
			sym := string(sym_rune)
			switch string(sym) {
			case ">", "<", "^", "v":
				carts[y][x] = append(carts[y][x], &Cart{string(sym), x, y, -1, -1, 0, &grid, &carts})
				if sym == ">" || sym == "<" {
					grid[y][x] = "-"
				} else if sym == "^" || sym == "v" {
					grid[y][x] = "|"
				}
			}
		}
	}
	//print(&grid, &carts)
	for {
		//fmt.Println("STEP:", i)
		//fmt.Println()
		step_all(&carts)
		//print(&grid, &carts)
	}
}

func step_all(carts_ptr *Carts) {
	// Useful to place to print the carts state
	// fmt.Println(*carts_ptr)
	var carts_to_be_stepped []*Cart
	for y, row := range *carts_ptr {
		Use(y)
		for _, col := range row {
			if len(col) == 1 {
				carts_to_be_stepped = append(carts_to_be_stepped, col[0])
			} else if len(col) > 1 {
				panic("CRASH!")
			}
		}
	}
	for _, cart_to_be_stepped_ptr := range carts_to_be_stepped {
		(*cart_to_be_stepped_ptr).step()
	}
}

func get_sym_beneath_cart(grid_ptr *[]string, cart Cart) {
	above_x := cart.x
	above_y := cart.y - 1
	above_sym := (*grid_ptr)[above_y][above_x]

	below_x := cart.x
	below_y := cart.y + 1
	below_sym := (*grid_ptr)[below_y][below_x]

	left_x := cart.x - 1
	left_y := cart.y
	left_sym := (*grid_ptr)[left_y][left_x]

	right_x := cart.x + 1
	right_y := cart.y
	right_sym := (*grid_ptr)[right_y][right_x]

	Use(above_sym, below_sym, left_sym, right_sym)
}

func (cart_ptr *Cart) get_next_coord() (next_x int, next_y int, next_sym string) {
	grid_ptr := cart_ptr.grid_ptr
	next_x = cart_ptr.x
	next_y = cart_ptr.y
	switch cart_ptr.sym {
	case ">":
		next_x += 1
	case "<":
		next_x -= 1
	case "^":
		next_y -= 1
	case "v":
		next_y += 1
	}
	next_sym = (*grid_ptr)[next_y][next_x]
	return
}

func array_delete(a *[]*Cart, i int) {
	// From https://github.com/golang/go/wiki/SliceTricks
	copy((*a)[i:], (*a)[i+1:])
	(*a)[len((*a))-1] = nil // or the zero value of T
	(*a) = (*a)[:len((*a))-1]
}
