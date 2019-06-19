package main

import (
	"bufio"
	"fmt"
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

type Square struct {
	id       int
	x        int
	y        int
	occupied bool
}
type Creature struct {
	id   int
	kind string
	x    int
	y    int
}
type Board struct {
	creatures        map[int]Creature
	squares          map[int]map[int]Square
	next_creature_id int
	next_square_id   int
}

func (board *Board) initialize() {
	board.squares = make(map[int]map[int]Square)
	board.creatures = make(map[int]Creature)
	board.next_creature_id = 1
	board.next_square_id = 1
}

func (board *Board) get_path(x1, y1, x2, y2 int) int{
	type DijkstraSquare struct {
		visited  bool
		distance int
		x        int
		y        int
		occupied bool
	}
	dsquares := make(map[int]map[int]*DijkstraSquare)
	var source, destination, current *DijkstraSquare
	for _, row := range board.squares {
		for _, square := range row {
			dsquare := DijkstraSquare{false, 1000000, square.x, square.y, square.occupied}
			if square.x == x1 && square.y == y1 {
				source = &dsquare
			}
			if square.x == x2 && square.y == y2 {
				destination = &dsquare
			}
			if dsquares[dsquare.y] == nil {
				dsquares[dsquare.y] = make(map[int]*DijkstraSquare)
			}
			dsquares[dsquare.y][dsquare.x] = &dsquare
		}
	}
	current = source
	(*current).distance = 0

	for {
		(*current).visited = true
		for _, neighbor := range board.get_neighbors(current.x, current.y) {
			dsquare_ptr := dsquares[neighbor.y][neighbor.x]
			(*dsquare_ptr).distance = current.distance + 1
		}

		new_current_ptr := &DijkstraSquare{false, 1000000, -1, -1, true}
		for _, row := range dsquares{
			for _, dsquare_ptr := range row{
				if dsquare_ptr.visited == false && dsquare_ptr.distance < new_current_ptr.distance{
					new_current_ptr = dsquare_ptr
				}
			}
		}
		current = new_current_ptr
		fmt.Println(current)
		if current.distance == 1000000 || dsquares[destination.y][destination.x].visited == true{
			break
		}
	}
	fmt.Println(dsquares)
	fmt.Println(dsquares[destination.y][destination.x])
	for _, row := range dsquares {
		for _, dsquare := range row{
			fmt.Print(dsquare.x, dsquare.y, dsquare.visited, " ")
		}
		fmt.Println()
	}
	return dsquares[destination.y][destination.x].distance
}

func (board *Board) get_neighbors(x, y int) (adjacent_squares []Square) {
	var square Square
	var present bool
	// left
	square, present = board.squares[y][x-1]
	if present {
		adjacent_squares = append(adjacent_squares, square)
	}
	// right
	square, present = board.squares[y][x+1]
	if present {
		adjacent_squares = append(adjacent_squares, square)
	}
	// up
	square, present = board.squares[y-1][x]
	if present {
		adjacent_squares = append(adjacent_squares, square)
	}
	// down
	square, present = board.squares[y+1][x-1]
	if present {
		adjacent_squares = append(adjacent_squares, square)
	}
	return
}

func (board *Board) get_all_of_kind(kind string) (creatures []Creature) {
	for _, creature := range board.creatures {
		if creature.kind == kind {
			creatures = append(creatures, creature)
		}
	}
	return
}

func (board *Board) move(id int, direction string) {
	creature := board.creatures[id]

	old_x := creature.x
	old_y := creature.y
	new_x := creature.x
	new_y := creature.y

	sq := board.squares[old_y][old_x]

	switch direction {
	case "up":
		new_y += 1
	case "down":
		new_y -= 1
	case "left":
		new_x -= 1
	case "right":
		new_x += 1
	}

	new_sq, present := board.squares[new_y][new_x]
	fmt.Println("new square", new_sq)
	if !present {
		panic("Attempting to move creature to a non-existent space")
	}
	if new_sq.occupied {
		panic("Attempting to move creature to an occupied space")
	}

	sq.occupied = false
	board.squares[old_y][old_x] = sq

	new_sq.occupied = true
	board.squares[new_y][new_x] = new_sq

	creature.x = new_x
	creature.y = new_y
	board.creatures[id] = creature
}

func (board *Board) insert(x, y int, sym string) {
	if sym == "#" {
		return
	}

	if board.squares[y] == nil {
		board.squares[y] = make(map[int]Square)
	}
	board.squares[y][x] = Square{board.next_square_id, x, y, false}
	board.next_square_id++

	if sym == "G" || sym == "E" {
		sq := board.squares[y][x]
		sq.occupied = true
		board.squares[y][x] = sq

		var creature_type string
		if sym == "G" {
			creature_type = "goblin"
		} else if sym == "E" {
			creature_type = "elf"
		}
		board.creatures[board.next_creature_id] = Creature{board.next_creature_id, creature_type, x, y}
		board.next_creature_id++
	}
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

	var board Board
	board.initialize()

	for y, row := range grid {
		for x, sym := range row {
			board.insert(x, y, sym)
		}
		fmt.Println(row)
	}

	elves := board.get_all_of_kind("elf")
	goblins := board.get_all_of_kind("goblin")
	fmt.Println(elves[0], goblins[2])
	fmt.Println(board.get_path(elves[0].x, elves[0].y, goblins[0].x, goblins[0].y))
}

func test_move(board *Board) {
	// Goblin #1, at 12, 2
	// move left, move right, move into a wall
	board.move(1, "left")
	fmt.Println(board.creatures[1])
	board.move(1, "right")
	fmt.Println(board.creatures[1])
	board.move(1, "right")
	fmt.Println(board.creatures[1])
}
