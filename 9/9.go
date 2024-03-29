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
	num_players := 476
	last_marble_worth := 71657

	// Test input
	// num_players := 13
	// last_marble_worth := 7999

	scores := make(map[int]int)

	// Bootstrap this thing
	var circle = []int{0, 1} // index is position, value is marble #
	current_player := 2
	current_index := 1
	current_marble := 2

	Use(circle, current_index, current_marble, last_marble_worth, scores)

	for current_marble <= last_marble_worth {
		var new_index int
		if mod(current_marble, 23) == 0 {
			// Add the current marble to player's score
			scores[current_player] = scores[current_player] + current_marble
			// Move the current_index 7 places to the left
			new_index = mod((current_index - 7), len(circle))
			if new_index < 0 {
				new_index += len(circle)
			}
			//fmt.Println(current_index, len(circle), new_index, new_index+len(circle))
			// Add the marble at that position the player's score
			scores[current_player] = scores[current_player] + circle[new_index]
			// Remove that marble
			circle = append(circle[:new_index], circle[new_index+1:]...)
		} else {
			var modval int
			if len(circle) == 0 {
				modval = 1
			} else if len(circle) == 1 {
				modval = 2
			} else if len(circle) == current_index+2 {
				modval = len(circle) + 1
			} else {
				modval = len(circle)
			}
			new_index = mod(current_index+2, modval)

			splice(&circle, current_marble, new_index)
		}

		//fmt.Println(current_player, circle)
		current_index = new_index
		current_marble++
		current_player = mod(current_player+1, num_players)
	}
	print_max_score(scores)
}

func print_max_score(scores map[int]int) {
	var max int
	var max_elf int
	for elf, score := range scores {
		if score > max {
			max = score
			max_elf = elf
		}
	}
	fmt.Println("Max score", max, "by elf", max_elf)
}

func mod(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func splice(a *[]int, x int, i int) {
	// https://github.com/golang/go/wiki/SliceTricks
	*a = append((*a)[:i], append([]int{x}, (*a)[i:]...)...)
}
