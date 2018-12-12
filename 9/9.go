package main

import (
	"fmt"
	"math"
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
	// Actual input: 476 players; last marble is worth 71657 points
	// num_players := 476
	// last_marble_worth := 71657

	// Test input
	num_players := 10
	last_marble_worth := 1618
	high_score := 8317

	var circle []int // index is position, value is marble #
	current_player := 0
	current_index := 0
	current_marble := 0

	Use(circle, current_index, current_marble, last_marble_worth, high_score)
	// rotate through the number of players
	// increment the marble #
	// update the slice
	// check if marble is multiple of 23

	n := 0
	for current_player < num_players && n < 10 {
		//fmt.Println(current_player)
		var modval int
		if len(circle) == 0 {
			modval = 1
		} else if len(circle) == 1 {
			modval = 2
		} else {
			modval = len(circle)
		}
		new_index := mod(current_index+2, modval)

		splice(&circle, current_marble, new_index)
		fmt.Println(circle)

		n++
		current_index = new_index
		current_marble++
		current_player = mod(current_player+1, num_players)
	}
}

func mod(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func splice(a *[]int, x int, i int) {
	// https://github.com/golang/go/wiki/SliceTricks
	*a = append((*a)[:i], append([]int{x}, (*a)[i:]...)...)
}
