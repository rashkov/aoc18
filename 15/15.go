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
	x int
	y int
}
type Creature struct {
	id   int
	kind string
	x int
	y int
}
type Board struct {
	creatures []Creature
	squares   []Square
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

	var id int
	for y, row := range grid {
		for x, sym := range row {
			switch sym {
			case "G":
				row[x] = "."
				board.creatures = append(board.creatures, Creature{id, "goblin", x, y})
				id++
			case "E":
				row[x] = "."
				board.creatures = append(board.creatures, Creature{id, "elf", x, y})
				id++
			case ".":
				board.squares = append(board.squares, Square{x,y})
			}
		}
		fmt.Println(row)
	}
	fmt.Println(board.creatures)
	fmt.Println(board.squares)
}
